{
    "name": "add_8bit",
    "nodes": {
        "add_4bit_0": "add_4bit",
        "add_4bit_1": "add_4bit"
    },
    "inputs": {
        "a_0": [
            {
                "uid": "add_4bit_0",
                "pid": "a_0"
            }
        ],
        "a_1": [
            {
                "uid": "add_4bit_0",
                "pid": "a_1"
            }
        ],
        "a_2": [
            {
                "uid": "add_4bit_0",
                "pid": "a_2"
            }
        ],
        "a_3": [
            {
                "uid": "add_4bit_0",
                "pid": "a_3"
            }
        ],
        "a_4": [
            {
                "uid": "add_4bit_1",
                "pid": "a_0"
            }
        ],
        "a_5": [
            {
                "uid": "add_4bit_1",
                "pid": "a_1"
            }
        ],
        "a_6": [
            {
                "uid": "add_4bit_1",
                "pid": "a_2"
            }
        ],
        "a_7": [
            {
                "uid": "add_4bit_1",
                "pid": "a_3"
            }
        ],
        "b_0": [
            {
                "uid": "add_4bit_0",
                "pid": "b_0"
            }
        ],
        "b_1": [
            {
                "uid": "add_4bit_0",
                "pid": "b_1"
            }
        ],
        "b_2": [
            {
                "uid": "add_4bit_0",
                "pid": "b_2"
            }
        ],
        "b_3": [
            {
                "uid": "add_4bit_0",
                "pid": "b_3"
            }
        ],
        "b_4": [
            {
                "uid": "add_4bit_1",
                "pid": "b_0"
            }
        ],
        "b_5": [
            {
                "uid": "add_4bit_1",
                "pid": "b_1"
            }
        ],
        "b_6": [
            {
                "uid": "add_4bit_1",
                "pid": "b_2"
            }
        ],
        "b_7": [
            {
                "uid": "add_4bit_1",
                "pid": "b_3"
            }
        ],
        "c_in": [
            {
                "uid": "add_4bit_0",
                "pid": "c_in"
            }
        ]
    },
    "edges": [
        {
            "from": {
                "uid": "add_4bit_0",
                "pid": "c_out"
            },
            "to": {
                "uid": "add_4bit_1",
                "pid": "c_in"
            }
        }
    ],
    "outputs": {
        "sum_0": {
            "uid": "add_4bit_0",
            "pid": "sum_0"
        },
        "sum_1": {
            "uid": "add_4bit_0",
            "pid": "sum_1"
        },
        "sum_2": {
            "uid": "add_4bit_0",
            "pid": "sum_2"
        },
        "sum_3": {
            "uid": "add_4bit_0",
            "pid": "sum_3"
        },
        "sum_4": {
            "uid": "add_4bit_1",
            "pid": "sum_0"
        },
        "sum_5": {
            "uid": "add_4bit_1",
            "pid": "sum_1"
        },
        "sum_6": {
            "uid": "add_4bit_1",
            "pid": "sum_2"
        },
        "sum_7": {
            "uid": "add_4bit_1",
            "pid": "sum_3"
        },
        "c_out": {
            "uid": "add_4bit_1",
            "pid": "c_out"
        }
    }
}