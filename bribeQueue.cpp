#include </Users/parvez/Private/cpp/std++.h>

using namespace std;

vector<string> split_string(string);

// // swap i with j
// vector<int> swap(vector<int> q, int i, int j){
//     int temp = q[i];
//     q[i] = q[j];
//     q[j] = temp;
//     return q;
// }

// // Complete the minimumBribes function below.
// void minimumBribes(vector<int> q) {
//     int bribe = 0;
//     // Traversing the current queue from last to 0
//     for(int i=q.size() - 1; i >=0 ; i--){
//         // Checking if ith element is in correct position
//         if(q[i] != (i + 1)){
//             if(((i-1) >= 0) && q[i - 1] == ( i + 1)){
//                 bribe += 1;
//                 q = swap(q, i-1, i);
//             } else if(((i-1) >= 0) && q[i - 2] == ( i + 1 )){
//                 bribe += 2;
//                 q = swap(q, i-2, i-1);
//                 q = swap(q, i-1, i);
//             } else{
//                 // More than 2 bribes not allowed
//                 bribe = INT_MIN;
//                 cout << "Too chaotic" << endl;
//                 break;
//             }
//         }
//     }
//     if(bribe != INT_MIN){
//         cout << bribe << endl;
//     }
// }


// Final accepted approach
// Complete the minimumBribes function below.
void minimumBribes(vector<int> q) {
    int bribe = 0;
    // Traversing the current queue from last to 0
    for(int i=q.size() - 1; i >=0 ; i--){
        // check if current position is more than 3
        if(q[i] - (i+1) > 2){
            bribe = INT_MIN;
            cout << "Too chaotic" << endl;
            break;
        }
        //Check how many numbers are bigger than the current i from it's original -2 position
        for(int j = max(0, q[i] - 2); j < i; j++){
            if(q[j] > q[i])
                bribe ++;
        }
    }
    if(bribe != INT_MIN){
        cout << bribe << endl;
    }
}


int main()
{
    int t;
    cin >> t;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    for (int t_itr = 0; t_itr < t; t_itr++) {
        int n;
        cin >> n;
        cin.ignore(numeric_limits<streamsize>::max(), '\n');

        string q_temp_temp;
        getline(cin, q_temp_temp);

        vector<string> q_temp = split_string(q_temp_temp);

        vector<int> q(n);

        for (int i = 0; i < n; i++) {
            int q_item = stoi(q_temp[i]);

            q[i] = q_item;
        }

        minimumBribes(q);
    }

    return 0;
}

vector<string> split_string(string input_string) {
    string::iterator new_end = unique(input_string.begin(), input_string.end(), [] (const char &x, const char &y) {
        return x == y and x == ' ';
    });

    input_string.erase(new_end, input_string.end());

    while (input_string[input_string.length() - 1] == ' ') {
        input_string.pop_back();
    }

    vector<string> splits;
    char delimiter = ' ';

    size_t i = 0;
    size_t pos = input_string.find(delimiter);

    while (pos != string::npos) {
        splits.push_back(input_string.substr(i, pos - i));

        i = pos + 1;
        pos = input_string.find(delimiter, i);
    }

    splits.push_back(input_string.substr(i, min(pos, input_string.length()) - i + 1));

    return splits;
}