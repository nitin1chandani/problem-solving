type RandomizedSet struct {
    set map[int]bool
}


func Constructor() RandomizedSet {
    return RandomizedSet{
        set: make(map[int]bool),
    }
}


func (this *RandomizedSet) Insert(val int) bool {
    if this.set[val] {
        return false
    }
    this.set[val] = true
    return true
}


func (this *RandomizedSet) Remove(val int) bool {
    if !this.set[val]{
        return false
    }
    delete(this.set, val)
    return true
}


func (this *RandomizedSet) GetRandom() int {
    keys := make([]int, 0, len(this.set))
    for key := range this.set{
        keys = append(keys, key)
    }
    return keys[rand.Intn(len(keys))]
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */