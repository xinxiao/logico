local pin = import '~pin.libsonnet';
local edge = import '~edge.libsonnet';

{
    "name": "if",
    "nodes": {
        "not": "not",
        "and_0": "and",
        "and_1": "and",
        "or": "or"
    },
    "inputs": {
        "a": [
            pin.Pin("and_0", "a"),
        ],
        "b": [
            pin.Pin("and_1", "a"),
        ],
        "cond": [
            pin.Pin("and_0", "b"),
            pin.Pin("not", "v"),
        ]
    },
    "edges": [
        edge.Edge(pin.Pin("not", "out"), pin.Pin("and_1", "b")),
        edge.Edge(pin.Pin("and_0", "out"), pin.Pin("or", "a")),
        edge.Edge(pin.Pin("and_1", "out"), pin.Pin("or", "b")),
    ],
    "outputs": {
        "out": pin.Pin("or", "out"),
    }
}