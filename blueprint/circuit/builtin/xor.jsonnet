local pin = import '~pin.libsonnet';
local edge = import '~edge.libsonnet';

{
    "name": "xor",
    "nodes": {
        "or": "or",
        "nand": "nand",
        "and": "and"
    },
    "inputs": {
        "a": [
            pin.Pin("or", "a"),
            pin.Pin("nand", "a"),
        ],
        "b": [
            pin.Pin("or", "b"),
            pin.Pin("nand", "b"),
        ]
    },
    "edges": [
        edge.Edge(pin.Pin("or", "out"), pin.Pin("and", "a")),
        edge.Edge(pin.Pin("nand", "out"), pin.Pin("and", "b")),
    ],
    "outputs": {
        "out": pin.Pin("and", "out")
    }
}