/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findPartTree(head *ListNode, root *TreeNode) bool {
    if head == nil {
        return true
    } else if root == nil{
        return false
    }
    
    return head.Val == root.Val && (findPartTree(head.Next, root.Left) || findPartTree(head.Next, root.Right))
}
func isSubPath(head *ListNode, root *TreeNode) bool {
    if root == nil {
        return false
    }
    return isSubPath(head, root.Left) || isSubPath(head, root.Right) || findPartTree(head, root)
}
