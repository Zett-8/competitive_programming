/**
 Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NULL
 */

/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @param {number} k
 * @return {ListNode}
 */

const rotateRight = (head, k) => {
  const listLen = listNode_len(head)
  k = k % listLen
  const n_fromHead = listLen - k

  let count = 1
  let current = head

  while (current && k) {
    if (count === n_fromHead) {
      const f = current.next
      current.next = null
      return connectLists(f, head)
    }

    current = current.next
    count++
  }

  return head
}

const connectLists = (first, second) => {
  let current = first

  while (current) {
    if (!current.next) {
      current.next = second
      return first
    }
    current = current.next
  }

  return second
}

const listNode_len = (head) => {
  let count = 0
  let current = head

  while (current) {
    count++
    current = current.next
  }

  return count
}
