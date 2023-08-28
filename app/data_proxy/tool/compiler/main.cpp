//go:build ignore
#include <iostream>
#include <vector>

using namespace std;

// 把大的往上调
void heapfity(vector<int> &v, int fa, int n) {
    int son = fa*2+1;
    while(son < n) {
        if(son<n-1 && v[son] < v[son+1]) son++;
        if(v[fa] < v[son]) {
            swap(v[fa], v[son]);
            fa = son, son = fa*2+1;
        } else {
            break;
        }
    }
}

void heapSort(vector<int> &v, int n) {
    // 调整堆，使最大的在堆顶
    // 从最后一个非叶子节点开始，把这个节点替换成子节点中最大的，再倒序执行，往上推
    for(int i = (n-1)/2; i >= 0; i--) {
        heapfity(v, i, n);
    }
    // 把最大值依次放到最后
    for(int i = n-1; i > 0; i--) {
        swap(v[i], v[0]);
        heapfity(v, 0, i); // 已经有一个排完序了
    }
}


int main() {
    vector<int> v;
    v.push_back(1);
    v.push_back(2);
    v.push_back(3);
    v.push_back(4);
    v.push_back(5);
    v.push_back(6);
    v.push_back(7);
    v.push_back(8);
    v.push_back(9);
    heapSort(v, v.size());
    for(auto val : v) {
        cout << val << ' ';
    }
    return 0;
}
