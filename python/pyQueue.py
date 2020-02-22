from multiprocessing import Queue

def main():
    q = Queue()
    for i in range(10):
        #print(i)
        q.put(i)
    print(q.get())
    print(q.get())
    print(q.get())
    print(q.get())
    print(q.get())
    q.close()
#    q.join_thread()

if __name__ == "__main__":
    main()
