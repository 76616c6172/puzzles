/* Sully's challenge

Here's a challenge - Take an input string parameter and determine if exactly 3 question marks exist between every pair of numbers that add up to 10. If so, return true, otherwise return false. Some examples are

TEST CASES:
"arrb6???4xxbl5???eee5" => true
"acc?7??sss?3rr1??????5" => true
"5??aaaaaaaaaaaaaaaaaaa?5?5" => false
"9???1???9???1???9" => true
"aa6?9" => false */

#import <stdio.h>
#import <ctype.h>
#import <stdbool.h>

/* transforms ASCII encoded numbers to integers */
int transASCIINum (int n) {
    return (n-48);
 }

int main () {
    bool flag = false; //true is 1, false is 0
    int c, cValue, lastNum, qmCount;
    c = cValue = lastNum = qmCount = 0;
    //int sum = 0;

    while ((c=getchar()) != EOF) {
        //if current char is a digit, transform the ascii encoding from getchar
        if (isdigit(c)) {
            cValue = transASCIINum(c);
            //printf("TRANSASCII REPORTING: %d\n",cValue);


            //the sum is 10 and the counter is NOT 3
            if (((cValue+lastNum) == 10) && (qmCount != 3)) {
                //Absolute fail, no need to do further checks.
                return false;
                //printf("ABSOLUTE FAIL SET\n");
            }
            //the sum is 10 and the counter IS 3
            if (((cValue+lastNum) == 10) && (qmCount == 3)) {
               //printf("ADDITION TO 10 DETECTED & COUNTER PASS ONCE\n");
               flag = true;
               qmCount = 0;
            }
            //the sum is NOT 10
            if ((cValue+lastNum) != 10) {
                qmCount = 0;
            }
        
            lastNum = cValue= transASCIINum(c);
            //printf("NUMBER DETECTED: %d\n",cValue);
        }
        else if ( c == '?') {
            qmCount++; 
            //printf("? DETECTED\n");
        }
    }

    if (flag == true ) {
        printf("TRUE\n");
        return true;
    }
//the entire string has been checked yet the flag is still false -> return FALSE
printf("FALSE\n");
return false;
}
