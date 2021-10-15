local pin = import '~pin.libsonnet';
local edge = import '~edge.libsonnet';

local BusPinID(bus, i) =
    bus + "_" + i;

{
    InputBus(name, size, input_pin_groups)::
        assert std.length(input_pin_groups) == size :
            "InputBus: number of input pin groups must match bus size";

        {   
            [BusPinID(name, i)]: input_pin_groups[i]
            for i in std.Range(size - 1)
        },
    

    OutputBus(name, size, output_pins)::
        assert std.length(output_pins) == size :
            "OutputBus: number of output pins must match bus size";

        {   
            [BusPinID(name, i)]: output_pins[i]
            for i in std.Range(size - 1)
        },
    

    BusToPinEdge(from_unit, from_bus, from_start, from_end, output_pins)::
        assert from_start >= 0 && from_end > from_start : 
            "BusEdge: invalid source bus endpoints";
        assert std.length(output_pins) == from_end - from_start :
            "BusEdge: number of output pins must match bus size";

        [
            edge.Edge(pin.Pin(from_unit, BusPinID(from_bus, i)), output_pins[i])
            for i in std.Range(from_start, from_end - 1)
        ],


    BusToBusEdge(from_unit, from_bus, from_start, from_end, to_unit, to_bus, to_start, to_end)::
        assert to_start >= 0 && to_end > to_start : 
            "BusEdge: invalid destination bus endpoints";
            
        self.BusToPinEdge(from_unit, from_bus, from_start, from_end, [
            pin.Pin(to_unit, BusPinID(to_bus, i))
            for i in std.Range(to_start, to_end) - 1
        ]),
}