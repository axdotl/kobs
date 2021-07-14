package resources

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/kobsio/kobs/pkg/api/clusters"
	"github.com/kobsio/kobs/pkg/api/middleware/errresponse"
	"github.com/kobsio/kobs/pkg/api/plugins/plugin"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"
)

// Route is the route under which the plugin should be registered in our router for the rest api.
const Route = "/resources"

var (
	log = logrus.WithFields(logrus.Fields{"package": "resources"})
)

// Resources is the structure for the getResources api call. It contains the cluster, namespace and the json
// representation of the retunred list object from the Kuberntes API.
type Resources struct {
	Cluster   string                 `json:"cluster"`
	Namespace string                 `json:"namespace"`
	Resources map[string]interface{} `json:"resources"`
}

// Config is the structure of the configuration for the resources plugin. It only contains one filed to forbid access to
// the provided resources.
type Config struct {
	Forbidden []string `json:"forbidden"`
}

// Router implements the router for the resources plugin, which can be registered in the router for our rest api.
type Router struct {
	*chi.Mux
	clusters *clusters.Clusters
	config   Config
}

// isForbidden checks if the requested resource was specified in the forbidden resources list. This can be used to use
// wildcard selectors in the RBAC permissions for kobs, but disallow the users to view secrets.
func (router *Router) isForbidden(resource string) bool {
	for _, r := range router.config.Forbidden {
		if resource == r {
			return true
		}
	}

	return false
}

// getResources returns a list of resources for the given clusters and namespaces. The result can limited by the
// paramName and param query parameter.
func (router *Router) getResources(w http.ResponseWriter, r *http.Request) {
	clusterNames := r.URL.Query()["cluster"]
	namespaces := r.URL.Query()["namespace"]
	resource := r.URL.Query().Get("resource")
	path := r.URL.Query().Get("path")
	paramName := r.URL.Query().Get("paramName")
	param := r.URL.Query().Get("param")

	log.WithFields(logrus.Fields{"clusters": clusterNames, "namespaces": namespaces, "resource": resource, "path": path, "paramName": paramName, "param": param}).Tracef("getResources")

	var resources []Resources

	// Loop through all the given cluster names and get for each provided name the cluster interface. After that we
	// check if the resource was provided via the forbidden resources list.
	for _, clusterName := range clusterNames {
		cluster := router.clusters.GetCluster(clusterName)
		if cluster == nil {
			render.Render(w, r, errresponse.Render(nil, http.StatusBadRequest, "invalid cluster name"))
			return
		}

		if router.isForbidden(resource) {
			render.Render(w, r, errresponse.Render(nil, http.StatusForbidden, fmt.Sprintf("access for resource %s is forbidding", resource)))
			return
		}

		// If the namespaces slice is nil, we retrieve the resource for all namespaces. If a list of namespaces was
		// provided we loop through all the namespaces and return the resources for these namespaces. All results are
		// added to the resources slice, which is then returned by the api.
		if namespaces == nil {
			list, err := cluster.GetResources(r.Context(), "", path, resource, paramName, param)
			if err != nil {
				render.Render(w, r, errresponse.Render(err, http.StatusBadRequest, "could not get resources"))
				return
			}

			var tmpResources map[string]interface{}
			err = json.Unmarshal(list, &tmpResources)
			if err != nil {
				render.Render(w, r, errresponse.Render(err, http.StatusInternalServerError, "could not unmarshal resources"))
				return
			}

			resources = append(resources, Resources{
				Cluster:   clusterName,
				Namespace: "",
				Resources: tmpResources,
			})
		} else {
			for _, namespace := range namespaces {
				list, err := cluster.GetResources(r.Context(), namespace, path, resource, paramName, param)
				if err != nil {
					render.Render(w, r, errresponse.Render(err, http.StatusBadRequest, "could not get resources"))
					return
				}

				var tmpResources map[string]interface{}
				err = json.Unmarshal(list, &tmpResources)
				if err != nil {
					render.Render(w, r, errresponse.Render(err, http.StatusInternalServerError, "could not unmarshal resources"))
					return
				}

				resources = append(resources, Resources{
					Cluster:   clusterName,
					Namespace: namespace,
					Resources: tmpResources,
				})
			}
		}
	}

	log.WithFields(logrus.Fields{"count": len(resources)}).Tracef("getResources")
	render.JSON(w, r, resources)
}

// getLogs returns the logs for the container of a pod in a cluster and namespace. A user can also set the time since
// when the logs should be returned.
func (router *Router) getLogs(w http.ResponseWriter, r *http.Request) {
	clusterName := r.URL.Query().Get("cluster")
	namespace := r.URL.Query().Get("namespace")
	name := r.URL.Query().Get("name")
	container := r.URL.Query().Get("container")
	since := r.URL.Query().Get("since")
	previous := r.URL.Query().Get("previous")

	log.WithFields(logrus.Fields{"cluster": clusterName, "namespace": namespace, "name": name, "container": container, "since": since, "previous": previous}).Tracef("getLogs")

	cluster := router.clusters.GetCluster(clusterName)
	if cluster == nil {
		render.Render(w, r, errresponse.Render(nil, http.StatusBadRequest, "invalid cluster name"))
		return
	}

	parsedSince, err := strconv.ParseInt(since, 10, 64)
	if err != nil {
		render.Render(w, r, errresponse.Render(err, http.StatusBadRequest, "could not parse since parameter"))
		return
	}

	parsedPrevious, err := strconv.ParseBool(previous)
	if err != nil {
		render.Render(w, r, errresponse.Render(err, http.StatusBadRequest, "could not parse previous parameter"))
		return
	}

	logs, err := cluster.GetLogs(r.Context(), namespace, name, container, parsedSince, parsedPrevious)
	if err != nil {
		render.Render(w, r, errresponse.Render(err, http.StatusBadGateway, "could not get logs"))
		return
	}

	log.WithFields(logrus.Fields{"count": len(logs)}).Tracef("getLogs")
	render.JSON(w, r, struct {
		Logs string `json:"logs"`
	}{logs})
}

// Register returns a new router which can be used in the router for the kobs rest api.
func Register(clusters *clusters.Clusters, plugins *plugin.Plugins, config Config) chi.Router {
	plugins.Append(plugin.Plugin{
		Name:        "resources",
		DisplayName: "Resources",
		Description: "View and edit Kubernetes resources.",
		Type:        "resources",
	})

	router := Router{
		chi.NewRouter(),
		clusters,
		config,
	}

	router.Get("/resources", router.getResources)
	router.Get("/logs", router.getLogs)

	return router
}