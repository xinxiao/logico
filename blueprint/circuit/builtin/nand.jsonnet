local pin = import '~pin.libsonnet';
local edge = import '~edge.libsonnet';

{
    "name": "nand",
    "nodes": {
        "and": "and",
        "not": "not"
    },
    "inputs": {
        "a": [
            pin.Pin("and", "a"),
        ],
        "b": [
            pin.Pin("and", "b"),
        ],
    },
    "edges": [
        edge.Edge(pin.Pin("and", "out"), pin.Pin("not", "v")),
    ],
    "outputs": {
        "out": pin.Pin("not", "out"),
    }
}