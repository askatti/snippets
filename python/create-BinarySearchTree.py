class BinarySearchTree:
    def __init__(self, value):
        self.value = value
        self.left_child = None
        self.right_child = None

    def insert_node(self, value):
        if value <= self.value and self.left_child:
            self.left_child.insert_node(value)
        elif value <= self.value:
            self.left_child = BinarySearchTree(value)
        elif value > self.value and self.right_child:
            self.right_child.insert_node(value)
        else:
            self.right_child = BinarySearchTree(value)
    def pre_order(self):
        """Traverse in pre-order self->left->right>."""
        if self.value:
            print( "%*s" % (15,self.value))
        if self.left_child:
            print( "%*s" % (-7,self.left_child.pre_order()))
        if self.right_child:
            print( "%*s" % (5,self.right_child.pre_order()))
 
    def find_node(self, root, value):
        if value == root.value or root == None: 
            print("value {} is present in {} node".format(value, root))
        elif value < root.value and root.left_child: 
            root.left_child.find_node(root.left_child, value)
        elif value > root.value and root.right_child:
            root.right_child.find_node( root.right_child, value)
        else:
            print("Num {} not found".format(value))

ar = [10,8,12,20,17,25,19]
print (ar)
bst = BinarySearchTree(15)
for i in ar:
    bst.insert_node(i)
bst.pre_order()
print("Enter num to serach")
input = int(input())
bst.find_node(bst, input)
