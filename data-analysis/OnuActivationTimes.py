import sys
import pandas as pd
import plotly.graph_objects as go
from plotly.offline import plot
import plotly.io as pio


# read input csv file
df = pd.read_csv(sys.argv[1])

print(df.head)
df.sort_values('ONUSTARTEVENT-ONU-discovery-indication-sent',inplace=True,ascending=True)
print(df.head)
# Plot the time diff EpochTime for ONUSerials
fig = go.Figure(go.Bar(x = df['ONUSerial'], y = df['EpochTime'], text = df['ONUSerial'],
                name='ONU Activation Times'))

title = "Report Generator- ONU Activation times for Events: " + sys.argv[3] + " - " + sys.argv[2]
xlabel = "ONU Serial Numbers"
ylabel = sys.argv[2] + " - " + sys.argv[3] + " (ms)"
fig.update_layout(title=title,
                 xaxis_title=xlabel,
                 yaxis_title=ylabel,
                 plot_bgcolor='rgb(230, 230,230)',
                 showlegend=True)

print(fig)

fig.show()
plot(fig, output_type="file", filename="ReportGenerator-3_ONUActivationTimes-sorted.html")
