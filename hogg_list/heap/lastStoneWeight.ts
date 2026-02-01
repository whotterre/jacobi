class MaxHeap {
    private heap: number[] = [];

    constructor(arr: number[]) {
        this.heap = arr;
        this.buildHeap();
    }

    private buildHeap() {
        for (let i = Math.floor(this.heap.length / 2) - 1; i >= 0; i--) {
            this.heapifyDown(i);
        }
    }

    private heapifyDown(i: number) {
        const n = this.heap.length;
        while (true) {
            let largest = i;
            let left = 2 * i + 1;
            let right = 2 * i + 2;

            if (left < n && this.heap[left] > this.heap[largest]) {
                largest = left;
            }
            if (right < n && this.heap[right] > this.heap[largest]) {
                largest = right;
            }

            if (largest !== i) {
                [this.heap[i], this.heap[largest]] = [this.heap[largest], this.heap[i]];
                i = largest;
            } else {
                break;
            }
        }
    }

    private heapifyUp(i: number) {
        while (i > 0) {
            let parent = Math.floor((i - 1) / 2);
            if (this.heap[parent] < this.heap[i]) {
                [this.heap[parent], this.heap[i]] = [this.heap[i], this.heap[parent]];
                i = parent;
            } else break;
        }
    }

    pop(): number | undefined {
        if (this.heap.length === 0) return undefined;
        if (this.heap.length === 1) return this.heap.pop();

        const max = this.heap[0];
        this.heap[0] = this.heap.pop()!;
        this.heapifyDown(0);
        return max;
    }

    push(val: number) {
        this.heap.push(val);
        this.heapifyUp(this.heap.length - 1);
    }

    size(): number {
        return this.heap.length;
    }
}

function lastStoneWeight(stones: number[]): number {
    const heap = new MaxHeap(stones);

    while (heap.size() > 1) {
        const y = heap.pop()!; // largest
        const x = heap.pop()!; // second largest

        if (y !== x) {
            heap.push(y - x);
        }
    }

    return heap.size() === 1 ? heap.pop()! : 0;
}
