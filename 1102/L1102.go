package main

import (
	"container/list"
	"fmt"
)

/*
1102. 得分最高的路径
给你一个 R 行 C 列的整数矩阵 A。矩阵上的路径从 [0,0] 开始，在 [R-1,C-1] 结束。
路径沿四个基本方向（上、下、左、右）展开，从一个已访问单元格移动到任一相邻的未访问单元格。
路径的得分是该路径上的 最小 值。例如，路径 8 →  4 →  5 →  9 的值为 4 。
找出所有路径中得分 最高 的那条路径，返回其 得分。

示例 1：
输入：[[5,4,5],[1,2,6],[7,4,6]]
输出：4
解释：
得分最高的路径用黄色突出显示。

示例 2：
输入：[[2,2,1,2,2,2],[1,2,2,2,1,2]]
输出：2

示例 3：
输入：[[3,4,6,3,4],[0,2,1,1,7],[8,8,3,2,7],[3,2,4,9,8],[4,1,2,0,0],[4,6,5,4,3]]
输出：3
 */

//type unionFind struct {
//	nums []int
//	count int
//}
//
//func Constructor1102(n int) unionFind {
//	nums := make([]int, n)
//	for i := 0; i < n; i++ {
//		nums[i] = i
//	}
//}
//
//func (u *unionFind) find(n int) int {
//	for u.nums[n] != n {
//		n = u.nums[n]
//	}
//	return n
//}
//
//func (u *unionFind) union(i, j int) {
//	root1 := u.find(i)
//	root2 := u.find(j)
//	if root1 != root2 {
//		if root1 < root2 {
//			u.nums[root2] = root1
//		} else {
//			u.nums[root1] = root2
//		}
//	}
//}

func min1102(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max1102(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type point struct {
	x int
	y int
}

var direction = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func maximumMinimumPath(A [][]int) int {
	row := len(A)
	column := len(A[0])
	output := 0
	low := 1
	high := min1102(A[0][0], A[row - 1][column - 1])
	for low <= high {
		visited := make([][]int, row)
		for i := 0; i < row; i++ {
			visited[i] = make([]int, column)
		}
		visited[0][0] = 1
		mid := (low + high) /2
		if findPathMatrix(A, visited, mid) {
			output = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return output
}

func findPathMatrix(A, visited [][]int, mid int) bool {
	path := list.New()
	path.PushBack(point{0, 0})
	for path.Len() != 0 {
		node := path.Front().Value.(point)
		for i := 0; i < 4; i++ {
			direct := direction[i]
			x := node.x + direct[0]
			y := node.y + direct[1]
			if validNode(x, y, A) && visited[x][y] != 1 && A[x][y] >= mid {
				visited[x][y] = 1
				path.PushBack(point{x, y})
				if x == len(A) - 1 && y == len(A[0]) - 1 {
					return true
				}
			}
		}
		path.Remove(path.Front())
	}
	return false
}

func validNode(x, y int, A [][]int) bool {
	return x >= 0 && x < len(A) && y >= 0 && y < len(A[0])
}

	func main() {
	fmt.Println(maximumMinimumPath([][]int{{5,0,3,5,4,1}, {3,5,1,1,2,5}, {3,5,5,5,4,0}, {2,0,3,0,5,5}, {1,4,5,0,0,5}}))
}

//func maximumMinimumPath(A [][]int) int {
//
//	low := 1
//	high := min1102(A[0][0], A[len(A)-1][len(A[0])-1])
//	result := 0
//	for low <= high {
//		visit := make([][]int, len(A))
//		for i := 0; i < len(visit); i++ {
//			visit[i] = make([]int, len(A[0]))
//		}
//		visit[0][0] = 1
//		mid := (low + high) / 2
//		if findPathMatrix(A, visit, mid) == true {
//			result = mid
//			low = mid + 1
//		} else {
//			high = mid - 1
//		}
//	}
//	return result
//}
//

/*
class UnionFind {
public:
	UnionFind(int n)
	{
		for (int i = 0; i < n; i++) {
			m_pre.emplace_back(i);
		}
	}

	~UnionFind()
	{
		m_pre.clear();
	}

	int GetFather(const int idx) {
		int r = idx;
		while (m_pre[r] != r) {
			r = m_pre[r];
		}

		int cur = idx;
		int upper;
		while (cur != r) {
			upper = m_pre[cur];
			m_pre[cur] = r;
			cur = upper;
		}

		return r;
	}

	void UninSet(const int a, const int b) {
		int af = GetFather(a);
		int bf = GetFather(b);

		if (af == bf) {
			return;
		}

		if (af > bf) {
			swap(af, bf);
		}

		m_pre[bf] = af;
	}
private:
	vector<int> m_pre;
};

class Solution {
public:
    struct Vertex {
        int i;
        int j;
        int val;
        bool operator<(const Vertex& rhs) const { return val < rhs.val; }
    };

    int maximumMinimumPath(vector<vector<int>>& A)
    {
        int rn = A.size();
        int cn = A[0].size();

        const vector<pair<int, int>> dirs = { {0,1},{0,-1},{-1,0},{1,0} };
        UnionFind uf(rn * cn);

        int maxVal = min(A[0][0], A[rn - 1][cn - 1]);
        priority_queue<Vertex> myQ;
        vector<vector<int>> coloredSet(rn, vector<int>(cn,0));

        for (int i = 0; i < rn; ++i) {
            for (int j = 0; j < cn; ++j) {
                myQ.push({ i,j,A[i][j] });
            }
        }

        while (uf.GetFather(0) != uf.GetFather(rn * cn - 1)) {
            Vertex node = myQ.top();
            myQ.pop();
            auto x = node.i;
            auto y = node.j;

            maxVal = min(node.val, maxVal);
            coloredSet[x][y] = 1;
            for (auto& dir : dirs) {
                int nx = x + dir.first;
                int ny = y + dir.second;

                if (nx < 0 || nx >= rn || ny < 0 || ny >= cn ||
                    coloredSet[nx][ny] == 0) {
                    continue;
                }
                uf.UninSet(x * cn + y, nx * cn + ny);
            }
        }

        return maxVal;
    }
};
 */