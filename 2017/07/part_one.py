class Node:
    name: str
    weight: int
    children: list[Node]
    parent: Node | None

    def __init__(self, name, weight):
        self.name = name
        self.weight = weight
        self.children = []
        self.parent = None


class NodeGraph:
    def __init__(self):
        self.orphans = {}
        self.unweighted = {}

    def add(self, node):
        if node.name in self.unweighted:
            old_node = self.unweighted[node.name]
            old_node.weight = node.weight
            old_node.children = node.children
            del self.unweighted[node.name]
        else:
            self.orphans[node.name] = node
        for child in node.children:
            if child.name in self.orphans:
                old_node = self.orphans[child.name]
                child.weight = old_node.weight
                child.children = old_node.children
                del self.orphans[child.name]
            else:
                self.unweighted[child.name] = child


def parse_node(node_str):
    components = node_str.split()
    name = components[0]
    weight = int(components[1].strip("()"))
    node = Node(name, weight)
    if len(components) > 3:
        children_names = [x.strip(",") for x in components[3:]]
        node.children = [Node(x, 0) for x in children_names]
        for child in node.children:
            child.parent = node
    return node


def read_graph(filename):
    nodes = NodeGraph()
    with open(filename) as tree_file:
        for line in tree_file:
            if line:
                node = parse_node(line.strip())
                nodes.add(node)
    return nodes


def bottom_node(graph):
    bottom_name = list(graph.orphans)[0]
    return graph.orphans[bottom_name]


def bottom_node_from_file(filename):
    graph = read_graph(filename)
    return bottom_node(graph)


if __name__ == "__main__":
    print(bottom_node_from_file("./2017/07/input.txt").name)
