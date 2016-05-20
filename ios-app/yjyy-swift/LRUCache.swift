// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import Foundation

class LRUNode<K, V> {
    var next: LRUNode?
    var previous: LRUNode?
    var key: K
    var value: V?

    init(key: K, value: V?) {
        self.key = key
        self.value = value
    }
}

class LRULinkedList<K, V> {

    var head: LRUNode<K, V>?
    var tail: LRUNode<K, V>?

    init() {}

    func addToHead(node: LRUNode<K, V>) {
        if self.head == nil  {
            self.head = node
            self.tail = node
        } else {
            let temp = self.head

            self.head?.previous = node
            self.head = node
            self.head?.next = temp
        }
    }

    func remove(node: LRUNode<K, V>) {
        if node === self.head {
            if self.head?.next != nil {
                self.head = self.head?.next
                self.head?.previous = nil
            } else {
                self.head = nil
                self.tail = nil
            }
        } else if node.next != nil {
            node.previous?.next = node.next
            node.next?.previous = node.previous
        } else {
            node.previous?.next = nil
            self.tail = node.previous
        }
    }

    func display() -> String {
        var description = ""
        var current = self.head

        while current != nil {
            description += "Key: \(current!.key) Value: \(current?.value) \n"

            current = current?.next
        }
        return description
    }
}

class LRUCache<K:Hashable, V>: CustomStringConvertible {

    let capacity: Int
    var length = 0

    private let queue: LRULinkedList<K, V>
    private var hashtable: [K: LRUNode<K, V>]

    // LRUCache, capacity is the number of elements to keep in the Cache.
    init(capacity: Int) {
        self.capacity = capacity

        self.queue = LRULinkedList()
        self.hashtable = [K: LRUNode<K, V>](minimumCapacity: self.capacity)
    }
	
	func Map() -> [K: LRUNode<K, V>] {
		return self.hashtable
	}

    subscript (key: K) -> V? {
        get {
            if let node = self.hashtable[key] {
                self.queue.remove(node)
                self.queue.addToHead(node)

                return node.value
            } else {
                return nil
            }
        }

        set(value) {
            if let node = self.hashtable[key] {
                node.value = value

                self.queue.remove(node)
                self.queue.addToHead(node)
            } else {
                let node = LRUNode(key: key, value: value)

                if self.length < capacity {
                    self.queue.addToHead(node)
                    self.hashtable[key] = node

                    self.length += 1
                } else {
                    hashtable.removeValueForKey(self.queue.tail!.key)
                    self.queue.tail = self.queue.tail?.previous

                    if let node = self.queue.tail {
                        node.next = nil
                    }

                    self.queue.addToHead(node)
                    self.hashtable[key] = node
                }
            }
        }
    }

    var description: String {
        return "LRUCache(\(self.length)) \n" + self.queue.display()
    }
}
