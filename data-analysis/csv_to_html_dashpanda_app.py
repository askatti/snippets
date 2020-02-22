#!/usr/bin/python3
"""CSV file and prints the HTML dropdown table converted and served on localhost:8050."""

import argparse
import dash
import dash_table
import dash_html_components as html
import dash_core_components as dcc
import pandas as pd
import json
from dash.dependencies import Input, Output, State

parser = argparse.ArgumentParser(description='A Tool to convert CSV HTML Dropdown table hosted on http://localhost:8050')
requiredNamed = parser.add_argument_group('required named arguments')
requiredNamed.add_argument('-f', '--file', help='Input csv file name', required=True, type=str, )

args = parser.parse_args()

df = pd.read_csv(
    args.file
    )

app = dash.Dash(__name__)

app.layout = html.Div(
    [
        dcc.Dropdown(
            id="dropdown",
            options=[{'label': name, 'value': name} for name in df['OLTID'].unique()],
            value="MTL",
        ),
        dcc.Store(id="selected-rows", storage_type="memory"),
        html.Div(id="tables-container"),
        html.Div(
            children=dash_table.DataTable(
                id="table",
                columns=[{"name": i, "id": i} for i in df.columns],
            )
        ),
    ]
)

@app.callback(
    Output("selected-rows", "data"),
    [Input("dropdown", "value")],
    [State("selected-rows", "data")],
)

def display_output(value, storage):
    """@Display output."""
    if value is not None:
        return {"selected_rows": df[df["OLTID"].str.contains(value)].to_json()}

@app.callback(
    Output("table", "data"),
    [Input("table", "sort_action"), Input("selected-rows", "data")],
)

def update_graph(sorting_settings, rows):
    """@Update Graph."""
    _df = pd.read_json(rows["selected_rows"])
    if sorting_settings is not None and len(sorting_settings):
        for setting in sorting_settings:
            _df.sort_values(
                by=setting["column_id"],
                ascending=(setting["direction"] == "asc"),
                inplace=True,
            )
        return _df.to_dict("rows")

    else:
        return _df.to_dict("rows")


if __name__ == "__main__":
    app.run_server(debug=True)
