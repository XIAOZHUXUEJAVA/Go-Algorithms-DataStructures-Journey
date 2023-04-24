package backtracking

var res [][]int
var track []int

func permute(nums []int) [][]int {
	res = [][]int{}
	track = []int{}
	used := make([]bool, len(nums))
	permutationHelper(nums, used)
	return res
}

func permutationHelper(nums []int, used []bool) {
	if len(track) == len(nums) {
		res = append(res, append([]int{}, track...))
		return
	}
	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		track = append(track, nums[i])
		used[i] = true
		permutationHelper(nums, used)
		track = track[:len(track)-1]
		used[i] = false
	}
}
