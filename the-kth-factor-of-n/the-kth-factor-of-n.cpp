class Solution {
public:
    int factors[32], pos = 0;
    void getFactors(int n) {
        // going up only to its square root
        int lmt = pow(n, 0.5);
        // dealing with perfect squares
        if (lmt * lmt == n) factors[pos++] = lmt;
        else lmt++;
        // checking all possible values
        for (int i = 1; i < lmt; i++) {
            if (n % i == 0) {
                factors[pos++] = i;
                factors[pos++] = n / i;
            }
        }
    }
    int kthFactor(int n, int k) {
        // getting all the factors
        getFactors(n);
        // returning if k is largers than the size of factors
        if (k > pos) return -1;
        // ordering them
        sort(begin(factors), begin(factors) + pos);
        return factors[k - 1];
    }
};