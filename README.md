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
