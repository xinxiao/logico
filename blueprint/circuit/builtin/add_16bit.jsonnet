{
    "name": "add_16bit",
    "nodes": {
        "add_8bit_0": "add_8bit",
        "add_8bit_1": "add_8bit"
    },
    "inputs": {
        "a_0": [
            {
                "uid": "add_8bit_0",
                "pid": "a_0"
            }
        ],
        "a_1": [
            {
                "uid": "add_8bit_0",
                "pid": "a_1"
            }
        ],
        "a_2": [
            {
                "uid": "add_8bit_0",
                "pid": "a_2"
            }
        ],
        "a_3": [
            {
                "uid": "add_8bit_0",
                "pid": "a_3"
            }
        ],
        "a_4": [
            {
                "uid": "add_8bit_0",
                "pid": "a_4"
            }
        ],
        "a_5": [
            {
                "uid": "add_8bit_0",
                "pid": "a_5"
            }
        ],
        "a_6": [
            {
                "uid": "add_8bit_0",
                "pid": "a_6"
            }
        ],
        "a_7": [
            {
                "uid": "add_8bit_0",
                "pid": "a_7"
            }
        ],
        "a_8": [
            {
                "uid": "add_8bit_1",
                "pid": "a_0"
            }
        ],
        "a_9": [
            {
                "uid": "add_8bit_1",
                "pid": "a_1"
            }
        ],
        "a_10": [
            {
                "uid": "add_8bit_1",
                "pid": "a_2"
            }
        ],
        "a_11": [
            {
                "uid": "add_8bit_1",
                "pid": "a_3"
            }
        ],
        "a_12": [
            {
                "uid": "add_8bit_1",
                "pid": "a_4"
            }
        ],
        "a_13": [
            {
                "uid": "add_8bit_1",
                "pid": "a_5"
            }
        ],
        "a_14": [
            {
                "uid": "add_8bit_1",
                "pid": "a_6"
            }
        ],
        "a_15": [
            {
                "uid": "add_8bit_1",
                "pid": "a_7"
            }
        ],
        "b_0": [
            {
                "uid": "add_8bit_0",
                "pid": "b_0"
            }
        ],
        "b_1": [
            {
                "uid": "add_8bit_0",
                "pid": "b_1"
            }
        ],
        "b_2": [
            {
                "uid": "add_8bit_0",
                "pid": "b_2"
            }
        ],
        "b_3": [
            {
                "uid": "add_8bit_0",
                "pid": "b_3"
            }
        ],
        "b_4": [
            {
                "uid": "add_8bit_0",
                "pid": "b_4"
            }
        ],
        "b_5": [
            {
                "uid": "add_8bit_0",
                "pid": "b_5"
            }
        ],
        "b_6": [
            {
                "uid": "add_8bit_0",
                "pid": "b_6"
            }
        ],
        "b_7": [
            {
                "uid": "add_8bit_0",
                "pid": "b_7"
            }
        ],
        "b_8": [
            {
                "uid": "add_8bit_1",
                "pid": "b_0"
            }
        ],
        "b_9": [
            {
                "uid": "add_8bit_1",
                "pid": "b_1"
            }
        ],
        "b_10": [
            {
                "uid": "add_8bit_1",
                "pid": "b_2"
            }
        ],
        "b_11": [
            {
                "uid": "add_8bit_1",
                "pid": "b_3"
            }
        ],
        "b_12": [
            {
                "uid": "add_8bit_1",
                "pid": "b_4"
            }
        ],
        "b_13": [
            {
                "uid": "add_8bit_1",
                "pid": "b_5"
            }
        ],
        "b_14": [
            {
                "uid": "add_8bit_1",
                "pid": "b_6"
            }
        ],
        "b_15": [
            {
                "uid": "add_8bit_1",
                "pid": "b_7"
            }
        ],
        "c_in": [
            {
                "uid": "add_8bit_0",
                "pid": "c_in"
            }
        ]
    },
    "edges": [
        {
            "from": {
                "uid": "add_8bit_0",
                "pid": "c_out"
            },
            "to": {
                "uid": "add_8bit_1",
                "pid": "c_in"
            }
        }
    ],
    "outputs": {
        "sum_0": {
            "uid": "add_8bit_0",
            "pid": "sum_0"
        },
        "sum_1": {
            "uid": "add_8bit_0",
            "pid": "sum_1"
        },
        "sum_2": {
            "uid": "add_8bit_0",
            "pid": "sum_2"
        },
        "sum_3": {
            "uid": "add_8bit_0",
            "pid": "sum_3"
        },
        "sum_4": {
            "uid": "add_8bit_0",
            "pid": "sum_4"
        },
        "sum_5": {
            "uid": "add_8bit_0",
            "pid": "sum_5"
        },
        "sum_6": {
            "uid": "add_8bit_0",
            "pid": "sum_6"
        },
        "sum_7": {
            "uid": "add_8bit_0",
            "pid": "sum_7"
        },
        "sum_8": {
            "uid": "add_8bit_1",
            "pid": "sum_0"
        },
        "sum_9": {
            "uid": "add_8bit_1",
            "pid": "sum_1"
        },
        "sum_10": {
            "uid": "add_8bit_1",
            "pid": "sum_2"
        },
        "sum_11": {
            "uid": "add_8bit_1",
            "pid": "sum_3"
        },
        "sum_12": {
            "uid": "add_8bit_1",
            "pid": "sum_4"
        },
        "sum_13": {
            "uid": "add_8bit_1",
            "pid": "sum_5"
        },
        "sum_14": {
            "uid": "add_8bit_1",
            "pid": "sum_6"
        },
        "sum_15": {
            "uid": "add_8bit_1",
            "pid": "sum_7"
        },
        "c_out": {
            "uid": "add_8bit_1",
            "pid": "c_out"
        }
    }
}