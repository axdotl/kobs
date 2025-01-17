import { Card, CardBody } from '@patternfly/react-core';
import { Datum, Node, ResponsiveScatterPlotCanvas, Serie } from '@nivo/scatterplot';
import React, { ReactNode, useMemo } from 'react';
import { SquareIcon } from '@patternfly/react-icons';
import { TooltipWrapper } from '@nivo/tooltip';
import Trace from './details/Trace';

import { ITrace } from '../../utils/interfaces';
import { doesTraceContainsError } from '../../utils/helpers';

interface IDatum extends Datum {
  label: string;
  size: number;
  trace: ITrace;
}

interface ITracesChartProps {
  name: string;
  traces: ITrace[];
  showDetails?: (details: React.ReactNode) => void;
}

function isIDatum(object: Datum): object is IDatum {
  return (object as IDatum).trace !== undefined;
}

const TracesChart: React.FunctionComponent<ITracesChartProps> = ({ name, traces, showDetails }: ITracesChartProps) => {
  const { series, min, max } = useMemo<{ series: Serie[]; min: number; max: number }>(() => {
    // Initialize min and max so that we can simply compare during traversing.
    let minimalSpans = Number.MAX_SAFE_INTEGER;
    let maximalSpans = 0;
    const result: IDatum[] = [];

    traces.forEach((trace) => {
      if (trace.spans.length < minimalSpans) {
        minimalSpans = trace.spans.length;
      }
      if (trace.spans.length > maximalSpans) {
        maximalSpans = trace.spans.length;
      }

      result.push({
        label: trace.traceName,
        size: trace.spans.length,
        trace,
        x: new Date(Math.floor(trace.spans[0].startTime / 1000)),
        y: trace.duration / 1000,
      });
    });

    return {
      max: maximalSpans,
      min: minimalSpans,
      series: [
        {
          data: result.sort((a, b) => (a.x.valueOf() as number) - (b.x.valueOf() as number)),
          id: 'Traces',
        },
      ],
    };
  }, [traces]);

  return (
    <Card isCompact={true}>
      <CardBody>
        <div style={{ height: '250px' }}>
          <ResponsiveScatterPlotCanvas
            axisBottom={{
              format: '%H:%M:%S',
            }}
            axisLeft={null}
            axisRight={null}
            axisTop={null}
            colors={['#0066cc']}
            data={series}
            enableGridX={false}
            enableGridY={false}
            margin={{ bottom: 25, left: 0, right: 0, top: 0 }}
            nodeSize={{ key: 'size', sizes: [15, 75], values: [min, max] }}
            theme={{
              background: '#ffffff',
              fontFamily: 'RedHatDisplay, Overpass, overpass, helvetica, arial, sans-serif',
              fontSize: 10,
              textColor: '#000000',
            }}
            onClick={(node: Node): void => {
              if (showDetails && isIDatum(node.data)) {
                showDetails(<Trace name={name} trace={node.data.trace} close={(): void => showDetails(undefined)} />);
              }
            }}
            // eslint-disable-next-line @typescript-eslint/ban-ts-comment
            // @ts-ignore as the typing expects NodeProps but actually has Node
            renderNode={(ctx: CanvasRenderingContext2D, props: Node): void => {
              // eslint-disable-next-line react/prop-types
              const hasError = isIDatum(props.data) ? doesTraceContainsError(props.data.trace) : false;

              ctx.beginPath();
              // eslint-disable-next-line react/prop-types
              ctx.arc(props.x, props.y, props.size / 2, 0, 2 * Math.PI);
              // eslint-disable-next-line react/prop-types
              ctx.fillStyle = hasError ? '#c9190b' : props.style.color;
              ctx.fill();
            }}
            tooltip={(tooltip): ReactNode => {
              const isFirstHalf = tooltip.node.index < series[0].data.length / 2;
              const hasError = isIDatum(tooltip.node.data) ? doesTraceContainsError(tooltip.node.data.trace) : false;
              const squareColor = hasError ? '#c9190b' : '#0066cc';

              return (
                <TooltipWrapper anchor={isFirstHalf ? 'right' : 'left'} position={[0, 20]}>
                  <div
                    className="pf-u-box-shadow-sm"
                    style={{
                      background: '#ffffff',
                      fontSize: '12px',
                      padding: '12px',
                      whiteSpace: 'nowrap',
                    }}
                  >
                    <div>
                      <b>{tooltip.node.data.formattedX}</b>
                    </div>
                    <div>
                      <SquareIcon color={squareColor} /> {(tooltip.node.data as unknown as IDatum).label}{' '}
                      {tooltip.node.data.formattedY}
                    </div>
                  </div>
                </TooltipWrapper>
              );
            }}
            xFormat="time:%Y-%m-%d %H:%M:%S.%L"
            xScale={{ type: 'time' }}
            yFormat={(e): string => e + ' ms'}
            yScale={{ max: 'auto', min: 'auto', type: 'linear' }}
          />
        </div>
      </CardBody>
    </Card>
  );
};

export default TracesChart;
