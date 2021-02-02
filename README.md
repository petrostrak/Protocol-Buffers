# Protocol-Buffers

### Advantages:
-Data is fully typed
-Data is compressed automatically (less CPU usage)
-Schema (defined using .proto file) is needed to generate code and read tht data
-Documentation can be embedded in the schema
-Data can be read acress any language
-Schema can evlove over time, in a safe manner (schema evolution)
-3-10x smaller, 20-100x faster than XML
-Code is generated for you automatically

### Disadvantages:
-Can't "open" the serialized data with a  text editor (because it's compressed and serialised)
