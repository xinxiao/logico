{
    "name": "negate_16bit",
    "nodes": {
        "flip_16bit": "flip_16bit",
        "add_16bit": "add_16bit"
    },
    "inputs": {
        "v_0": [
            {
                "uid": "flip_16bit",
                "pid": "v_0"
            }
        ],
        "v_1": [
            {
                "uid": "flip_16bit",
                "pid": "v_1"
            }
        ],
        "v_2": [
            {
                "uid": "flip_16bit",
                "pid": "v_2"
            }
        ],
        "v_3": [
            {
                "uid": "flip_16bit",
                "pid": "v_3"
            }
        ],
        "v_4": [
            {
                "uid": "flip_16bit",
                "pid": "v_4"
            }
        ],
        "v_5": [
            {
                "uid": "flip_16bit",
                "pid": "v_5"
            }
        ],
        "v_6": [
            {
                "uid": "flip_16bit",
                "pid": "v_6"
            }
        ],
        "v_7": [
            {
                "uid": "flip_16bit",
                "pid": "v_7"
            }
        ],
        "v_8": [
            {
                "uid": "flip_16bit",
                "pid": "v_8"
            }
        ],
        "v_9": [
            {
                "uid": "flip_16bit",
                "pid": "v_9"
            }
        ],
        "v_10": [
            {
                "uid": "flip_16bit",
                "pid": "v_10"
            }
        ],
        "v_11": [
            {
                "uid": "flip_16bit",
                "pid": "v_11"
            }
        ],
        "v_12": [
            {
                "uid": "flip_16bit",
                "pid": "v_12"
            }
        ],
        "v_13": [
            {
                "uid": "flip_16bit",
                "pid": "v_13"
            }
        ],
        "v_14": [
            {
                "uid": "flip_16bit",
                "pid": "v_14"
            }
        ],
        "v_15": [
            {
                "uid": "flip_16bit",
                "pid": "v_15"
            }
        ]
    },
    "always_on": [
        {
            "uid": "add_16bit",
            "pid": "b_0"
        }
    ],
    "always_off": [
        {
            "uid": "add_16bit",
            "pid": "c_in"
        },
        {
            "uid": "add_16bit",
            "pid": "b_1"
        },
        {
            "uid": "add_16bit",
            "pid": "b_2"
        },
        {
            "uid": "add_16bit",
            "pid": "b_3"
        },
        {
            "uid": "add_16bit",
            "pid": "b_4"
        },
        {
            "uid": "add_16bit",
            "pid": "b_5"
        },
        {
            "uid": "add_16bit",
            "pid": "b_6"
        },
        {
            "uid": "add_16bit",
            "pid": "b_7"
        },
        {
            "uid": "add_16bit",
            "pid": "b_8"
        },
        {
            "uid": "add_16bit",
            "pid": "b_9"
        },
        {
            "uid": "add_16bit",
            "pid": "b_10"
        },
        {
            "uid": "add_16bit",
            "pid": "b_11"
        },
        {
            "uid": "add_16bit",
            "pid": "b_12"
        },
        {
            "uid": "add_16bit",
            "pid": "b_13"
        },
        {
            "uid": "add_16bit",
            "pid": "b_14"
        },
        {
            "uid": "add_16bit",
            "pid": "b_15"
        }
    ],
    "edges": [
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_0"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_0"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_1"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_1"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_2"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_2"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_3"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_3"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_4"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_4"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_5"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_5"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_6"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_6"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_7"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_7"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_8"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_8"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_9"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_9"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_10"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_10"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_11"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_11"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_12"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_12"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_13"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_13"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_14"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_14"
            }
        },
        {
            "from": {
                "uid": "flip_16bit",
                "pid": "out_15"
            },
            "to": {
                "uid": "add_16bit",
                "pid": "a_15"
            }
        }
    ],
    "outputs": {
        "out_0": {
            "uid": "add_16bit",
            "pid": "sum_0"
        },
        "out_1": {
            "uid": "add_16bit",
            "pid": "sum_1"
        },
        "out_2": {
            "uid": "add_16bit",
            "pid": "sum_2"
        },
        "out_3": {
            "uid": "add_16bit",
            "pid": "sum_3"
        },
        "out_4": {
            "uid": "add_16bit",
            "pid": "sum_4"
        },
        "out_5": {
            "uid": "add_16bit",
            "pid": "sum_5"
        },
        "out_6": {
            "uid": "add_16bit",
            "pid": "sum_6"
        },
        "out_7": {
            "uid": "add_16bit",
            "pid": "sum_7"
        },
        "out_8": {
            "uid": "add_16bit",
            "pid": "sum_8"
        },
        "out_9": {
            "uid": "add_16bit",
            "pid": "sum_9"
        },
        "out_10": {
            "uid": "add_16bit",
            "pid": "sum_10"
        },
        "out_11": {
            "uid": "add_16bit",
            "pid": "sum_11"
        },
        "out_12": {
            "uid": "add_16bit",
            "pid": "sum_12"
        },
        "out_13": {
            "uid": "add_16bit",
            "pid": "sum_13"
        },
        "out_14": {
            "uid": "add_16bit",
            "pid": "sum_14"
        },
        "out_15": {
            "uid": "add_16bit",
            "pid": "sum_15"
        }
    }
}