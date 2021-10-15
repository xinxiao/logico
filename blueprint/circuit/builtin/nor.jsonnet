local pin = import '~pin.libsonnet';
local edge = import '~edge.libsonnet';

{
    "name": "nor",
    "nodes": {
        "or": "or",
        "not": "not"
    },
    "inputs": {
        "a": [
            pin.Pin("or", "a"),
        ],
        "b": [
            pin.Pin("or", "b"),
        ],
    },
    "edges": [
        edge.Edge(pin.Pin("or", "out"), pin.Pin("not", "v")),
    ],
    "outputs": {
        "out": pin.Pin("not", "out"),
    }
}