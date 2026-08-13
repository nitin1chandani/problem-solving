class Trie {

    Trie[] child;
    boolean eow;

    public Trie() {
        child = new Trie[26];
        eow = false;
    }
    
    public void insert(String word) {
        Trie curr = this;

        for(char c : word.toCharArray()){
            int idx = c - 'a';
            if(curr.child[idx]==null){
                curr.child[idx] = new Trie();
            }

            curr = curr.child[idx];
        }

        curr.eow = true;
    }
    
    public boolean search(String word) {
        Trie curr = this;

        for(char c : word.toCharArray()){
            int idx = c - 'a';

            if(curr.child[idx]==null){
                return false;
            }else{
                curr = curr.child[idx];
            }
        }

        return curr.eow ? true : false;
    }
    
    public boolean startsWith(String prefix) {
        Trie curr = this;

        for(char c : prefix.toCharArray()){
            int idx = c - 'a';
            if(curr.child[idx]==null){
                return false;
            }else{
                curr = curr.child[idx];
            }
        }

        return true;
    }
}

/**
 * Your Trie object will be instantiated and called as such:
 * Trie obj = new Trie();
 * obj.insert(word);
 * boolean param_2 = obj.search(word);
 * boolean param_3 = obj.startsWith(prefix);
 */