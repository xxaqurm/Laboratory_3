#pragma once
#include <string>

using namespace std;

template <typename DS>
void runInteractive(const string& name) {
    DS ds;

    while (true) {
        cout << "> ";
        string cmd;
        cin >> cmd;

        if (cmd == "add") {
            int x; 
            cin >> x;
            ds.push_back(x);
            cout << "< OK\n";
        }
        else if (cmd == "remove") {
            int x;
            cin >> x;
            ds.remove(x);
            cout << "< OK\n";
        }
        else if (cmd == "find") {
            int x;
            cin >> x;
            cout << (ds.contains(x) ? "< Exists\n" : "< Not found\n");
        }
        else if (cmd == "print") {
            cout << "< ";
            ds.display();
        }
        else if (cmd == "clear") {
            ds.clear();
            cout << "< Cleared\n";
        }
        else if (cmd == "exit") {
            break;
        }
        else {
            cout << "< Unknown command\n";
        }
    }
}

template <typename HashMap>
void runInteractiveHash(const string& name) {
    HashMap map;

    while (true) {
        cout << "> ";
        string cmd;
        cin >> cmd;

        if (cmd == "add") {
            int k, v;
            cin >> k >> v;
            map.put(k, v);
            cout << "< OK\n";
        }
        else if (cmd == "remove") {
            int k;
            cin >> k;
            map.remove(k);
            cout << "< OK\n";
        }
        else if (cmd == "find") {
            int k;
            cin >> k;
            cout << (map.contains(k) ? "< Exists\n" : "< Not found\n");
        }
        else if (cmd == "print") {
            cout << "< ";
            map.display();
        }
        else if (cmd == "exit") {
            break;
        }
        else {
            cout << "< Unknown command\n";
        }
    }
}
