{
    "name": "negate_8bit",
    "nodes": {
        "flip_8bit": "flip_8bit",
        "add_8bit": "add_8bit"
    },
    "inputs": {
        "v_0": [
            {
                "uid": "flip_8bit",
                "pid": "v_0"
            }
        ],
        "v_1": [
            {
                "uid": "flip_8bit",
                "pid": "v_1"
            }
        ],
        "v_2": [
            {
                "uid": "flip_8bit",
                "pid": "v_2"
            }
        ],
        "v_3": [
            {
                "uid": "flip_8bit",
                "pid": "v_3"
            }
        ],
        "v_4": [
            {
                "uid": "flip_8bit",
                "pid": "v_4"
            }
        ],
        "v_5": [
            {
                "uid": "flip_8bit",
                "pid": "v_5"
            }
        ],
        "v_6": [
            {
                "uid": "flip_8bit",
                "pid": "v_6"
            }
        ],
        "v_7": [
            {
                "uid": "flip_8bit",
                "pid": "v_7"
            }
        ]
    },
    "always_on": [
        {
            "uid": "add_8bit",
            "pid": "b_0"
        }
    ],
    "always_off": [
        {
            "uid": "add_8bit",
            "pid": "c_in"
        },
        {
            "uid": "add_8bit",
            "pid": "b_1"
        },
        {
            "uid": "add_8bit",
            "pid": "b_2"
        },
        {
            "uid": "add_8bit",
            "pid": "b_3"
        },
        {
            "uid": "add_8bit",
            "pid": "b_4"
        },
        {
            "uid": "add_8bit",
            "pid": "b_5"
        },
        {
            "uid": "add_8bit",
            "pid": "b_6"
        },
        {
            "uid": "add_8bit",
            "pid": "b_7"
        }
    ],
    "edges": [
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_0"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_0"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_1"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_1"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_2"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_2"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_3"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_3"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_4"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_4"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_5"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_5"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_6"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_6"
            }
        },
        {
            "from": {
                "uid": "flip_8bit",
                "pid": "out_7"
            },
            "to": {
                "uid": "add_8bit",
                "pid": "a_7"
            }
        }
    ],
    "outputs": {
        "out_0": {
            "uid": "add_8bit",
            "pid": "sum_0"
        },
        "out_1": {
            "uid": "add_8bit",
            "pid": "sum_1"
        },
        "out_2": {
            "uid": "add_8bit",
            "pid": "sum_2"
        },
        "out_3": {
            "uid": "add_8bit",
            "pid": "sum_3"
        },
        "out_4": {
            "uid": "add_8bit",
            "pid": "sum_4"
        },
        "out_5": {
            "uid": "add_8bit",
            "pid": "sum_5"
        },
        "out_6": {
            "uid": "add_8bit",
            "pid": "sum_6"
        },
        "out_7": {
            "uid": "add_8bit",
            "pid": "sum_7"
        }
    }
}