#include <stdio.h>

void merge(int a[], int m, int b[],int n) {

    while ( m > 0 && n > 0) {
        if (a[m-1] > b[n-1]) {
            a[m+n -1] = a[m -1];
            m--;
        } else {
            a[m+n-1] = b[n-1];
            n--;
        }
    }
    while ( n >0 ) {
        a[m+n-1] = b[n-1];
        n--;
    }
}
int main () {
    int a[] = {6,7,8,9};
    int b[] = {2,3,4,5};

    merge(a,4,b,4);
    for(int i=0; i< 8;i++)
        printf("%d\n",a[i]);
}
