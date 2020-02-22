from  multiprocessing import Queue
class binaryTree():
    """Sample Program to Learn Binary Tree in python."""
    def __init__(self, value):
        """Init basic Tree."""
        self.value = value
        self.left_child = None
        self.right_child = None

    def insert_left(self,value):
        """Insert node at left."""
        if self.left_child == None:
            self.left_child = binaryTree(value)
        else:
            node = binaryTree(value)
            node.left_child = self.left_child
            self.left_child = node

    def insert_right(self,value):
        """Insert node at right."""
        if self.right_child == None:
            self.right_child = binaryTree(value)
        else:
            node = binaryTree(value)
            node.right_child = self.right_child
            self.right_child = node

    def pre_order(self):
        """Traverse in pre-order self->left->right>."""
        if self.value:
            print(self.value),
        if self.left_child:
            print(self.left_child.pre_order()),
        if self.right_child:
            print(self.right_child.pre_order()),
 
    def in_order(self):
        """Inorder traverse."""
        if self.left_child:
            print(self.left_child.in_order()),
        if self.value != None:
            print(self.value),
        if self.right_child:
            print(self.right_child.in_order()),

    def post_order(self):
        """Inorder traverse."""
        if self.left_child:
            print(self.left_child.post_order()),
        if self.right_child:
            print(self.right_child.post_order()),
        if self.value != None:
            print(self.value),

    def bfs(self):
        global queue
        queue = Queue()
        queue.put(self)

        while not queue.empty():
            current_node = queue.get()
            print(current_node.value)

            if current_node.left_child:
                queue.put(current_node.left_child)

            if current_node.right_child:
                queue.put(current_node.right_child)
         
# Create this tree
#        1
#       / \
#      2   5
#     / \  /\
#    3   46  7
#
TreeA =  binaryTree("1")
TreeA.insert_left("2")
TreeA.insert_right("5")
b_node = TreeA.left_child
b_node.insert_left("3")
b_node.insert_right("4")
e_node = TreeA.right_child
e_node.insert_left("6")
e_node.insert_right("7")
print("%10s" % TreeA.value)
print('%*s %*s' % (8, "/", 3, "\\"))
print('%*s %*s' % (7, b_node.value, 5, e_node.value))
print('%*s %*s %*s %*s' % (6, "/", 1, "\\", 3, "/",1,"\\"))
print('%*s %*s %s %*s' % (5, b_node.left_child.value, 3, b_node.right_child.value,e_node.left_child.value,3,e_node.right_child.value))

#print("\nTraverse pre-order DFS")
#TreeA.pre_order()
#print("\nTraverse in-order DFS")
#TreeA.in_order()
#print("\nTraverse post_order DFS")
#TreeA.post_order()
print("\nTraverse BFS")
TreeA.bfs()
