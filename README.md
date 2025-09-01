# AEL

**Aether Execution Layer**<br>
A Command line interface for [Vanadium Aether&trade;](https://github.com/Vanadium-Development/aether) written in Go.

# Command Reference

---

### Configuration Management

An AEL Config encapsulates the entire AEL setup, including nodes, tracked
files, etc.
> **Initialize an AEL Configuration**<br>
> `ael init`

### Node Management

Nodes are physical computers running **Vanadium Aether&trade;** instances that
execute rendering jobs around AEL instructions.

> **List Nodes in Pool**<br>
> `ael node list`

> **Add Node to Pool**<br>
> `ael node add [--address | -a] <ip> [--name | -n] <name>`

### Scene Management

A Scene is a set of all required files to render a blender scene, consisting both of the blend file itself and
all other required resources.

> **Add files to current scene**<br>
> `ael scene add [--file | -f] <file path>`

> **Remove files from current scene**<br>
> `ael scene remove [--file | -f] <file path>`

> **Commit scene files to all registered nodes**<br>
> `ael commit`

## Rendering

> **Initialize distributed rendering process** <br>
> The rendering process of the previously commited scene will start on all registered nodes simultaneously.
> The rendering results will be collected from all nodes and stored in the specified directory.<br>
> `ael render [--output | -o] <directory>`