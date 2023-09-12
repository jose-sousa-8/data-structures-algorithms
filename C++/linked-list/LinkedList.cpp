#include <iostream>

using namespace std;

class Node { 
    public:
        int value;
        Node* next;

        Node(int value) {
            this->value = value;
            next = nullptr;
        }
}; 


class LinkedList {
    private:
        Node* head;
        Node* tail;
        int length;

    public:
        LinkedList(int value) {
            Node* newNode = new Node(value);
            head = newNode;
            tail = newNode;
            length = 1;
        }

        ~LinkedList() {
            Node* temp = head;
            while (head) {
                head = head->next;
                delete temp;
                temp = head;
            }
        }

        void printList() {
            Node* temp = head;
            if (temp == nullptr) {
                cout << "empty";
            } else {
                while (temp != nullptr) {
                    cout << temp->value;
                    temp = temp->next;
                    if (temp != nullptr) {
                        cout << " -> ";
                    }
                }
            }
            cout << endl;
        }

        Node* getHead() {
            return head;
        }

        Node* getTail() {
            return tail; 
        }

        int getLength() {
            return length;
        }
        
        // We will use this method to test append to an empty list
        void makeEmpty() {
            Node* temp = head;
            while (head) {
                head = head->next;
                delete temp;
                temp = head;
            }
            tail = nullptr;
            length = 0;
        }

        void append(int value) {
            Node* newNode = new Node(value);
            if (length == 0) {
                head = newNode;
                tail = newNode;
            } else {
                tail->next = newNode;
                tail = newNode;
            }
            length++;
        }

        void deleteLast() {
            if (length == 0) return;
            Node* temp = head;
            if (length == 1) {
                head = nullptr;
                tail = nullptr;
            } else {
                Node* pre = head;
                while(temp->next) {
                    pre = temp;
                    temp = temp->next;
                }
                tail = pre;
                tail->next = nullptr;
            }
            delete temp;
            length--;            
        }

        void prepend(int value) {
            Node* newNode = new Node(value);
            if (length == 0) {
                head = newNode;
                tail = newNode;
            } else {
                newNode->next = head;
                head = newNode;
            }
            length++;
        }

        void deleteFirst() {
            if (length == 0) return;
            Node* temp = head;
            if (length == 1) {
                head = nullptr;
                tail = nullptr;
            } else {
                head = head->next;
            }
            delete temp;
            length--;
        }

        Node* get(int index) {
            if (index < 0 || index >= length) return nullptr;
            Node* temp = head;
            for (int i = 0; i < index; ++i) {
                temp = temp->next;
            }
            return temp;
        }

        bool set(int index, int value) {
            Node* temp = get(index);
            if (temp) {
                temp->value = value;
                return true;
            } 
            return false;
        }

        bool insert(int index, int value) {
            if (index < 0 || index > length) return false;
            if (index == 0) {
                prepend(value);
                return true;
            }
            if (index == length) {
                append(value);
                return true;
            }
            Node* newNode = new Node(value);
            Node* temp = get(index - 1);
            newNode->next = temp->next;
            temp->next = newNode;
            length++;
            return true;
        }   

        void deleteNode(int index) {
            if (index < 0 || index >= length) return;
            if (index == 0) return deleteFirst();
            if (index == length - 1) return deleteLast();

            Node* prev = get(index - 1);
            Node* temp = prev->next;

            prev->next = temp->next;
            delete temp;
            length--;        
        }
        
        void reverse() {
            Node* temp = head;
            head = tail;
            tail = temp;
            Node* after = temp->next;
            Node* before = nullptr;
            for (int i = 0; i < length; ++i) {
            // Instead of the for loop you could use: 
            // while (temp != nullptr) {
            // -- or --
            // while (temp) {
                after = temp->next;
                temp->next = before;
                before = temp;
                temp = after;
            }
        }

        Node* findMiddleNode() {
            Node* slow = head;
            Node* fast = head;
            while (fast != nullptr && fast->next != nullptr) {
                slow = slow->next;
                fast = fast->next->next;
            }
            return slow;
        }

        bool hasLoop() {
            Node* slow = head;
            Node* fast = head;
            while (fast != nullptr && fast->next != nullptr) {
                slow = slow->next;
                fast = fast->next->next;
                if (slow == fast) {
                    return true;
                }
            }
            return false;
        }

        Node* findKthFromEnd(int k) {
            Node* slow = head;
            Node* fast = head;
            for (int i = 0; i < k; ++i) {
                if (fast == nullptr) {
                    return nullptr;
                }
                fast = fast->next;
            }
            while (fast != nullptr) {
                slow = slow->next;
                fast = fast->next;
            }
            return slow;
        }

        void removeDuplicates() {
            unordered_set<int> values;
            Node* previous = nullptr;
            Node* current = head;
            while (current != nullptr) {
                if (values.find(current->value) != values.end()) {
                    previous->next = current->next;
                    delete current;
                    current = previous->next;
                    length -= 1;
                } else {
                    values.insert(current->value);
                    previous = current;
                    current = current->next;
                }
            }
        }

        int binaryToDecimal() {
            int num = 0;
            Node* current = head;
            while (current != nullptr) {
                num = num * 2 + current->value;
                current = current->next;
            }
            return num;
        }

        void reverseBetween(int m, int n) {
            if (head == nullptr) return;
    
            Node* dummy = new Node(0);
            dummy->next = head;
            Node* prev = dummy;
    
            for (int i = 0; i < m; i++) {
                prev = prev->next;
            }
    
            Node* current = prev->next;
            for (int i = 0; i < n - m; i++) {
                Node* temp = current->next;
                current->next = temp->next;
                temp->next = prev->next;
                prev->next = temp;
            }
    
            head = dummy->next;
            delete dummy;
        }
}; 


