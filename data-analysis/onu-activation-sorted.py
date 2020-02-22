import sys
import pandas as pd
import plotly.io as pio

from plotly.offline import plot

# read input csv file
df = pd.read_csv(sys.argv[1])

aggs = ["avg","median","mode","rms","stddev","min","max","first","last"]

agg = []
agg_func = []
for i in range(0, len(aggs)):
    agg = dict(
        args=['transforms[0].aggregations[0].func', aggs[i]],
        label=aggs[i],
        method='restyle'
    )
    agg_func.append(agg)

data = [dict(
    type = 'scatter',
    x = df['ONUSerial'],
    y = df['EpochTime'],
    mode = 'markers',
    transforms = [dict(
        type = 'aggregate',
        groups = df['ONUSerial'],
        aggregations = [dict(
            target = 'y', func = 'avg', enabled = True)
        ]
    )]
)]

title = '<b>Report Generator- ONU Activation times for Events: " + sys.argv[3] + " - " + sys.argv[2] </b><br>use dropdown to change aggregation'
xlabel = "ONU Serial Numbers"
ylabel = sys.argv[2] + " - " + sys.argv[3] + " (ms)"

layout = dict(
    title = '<b>Plotly Aggregations</b><br>use dropdown to change aggregation',
    xaxis = dict(title = xlabel),
    yaxis = dict(title = ylabel),
    updatemenus = [dict(
        x = 0.85,
        y = 1.15,
        xref = 'paper',
        yref = 'paper',
        yanchor = 'top',
        active = 1,
        showactive = False,
        buttons = agg_func
    )]
)

fig_dict = dict(data=data, layout=layout)

pio.show(fig_dict, validate=False)
#pio.write_html(fig_dict,file=mytest-dropdown.html)

#plot(fig_dict, output_type="file", filename="ReportGenerator-3_ONUActivationTimes-sorted-dropdown.html")

