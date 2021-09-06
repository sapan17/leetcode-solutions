class Solution {
public:
int kthFactor(int n, int k) {

    vector<int>v;
    map<int,int>mp;
    
    for(int i=1;i<=sqrt(n);i++)
    {
       if(n%i==0)
       {
           if(mp[i]==0)
           {
               v.push_back(i);
               mp[i]=1;
           }
           if(mp[n/i]==0)
           {
               mp[n/i]=1;
               v.push_back(n/i);
           }
       }
    }
    sort(v.begin(),v.end());
    if(v.size()<k)
        return -1;
    return v[k-1];
    
}
};