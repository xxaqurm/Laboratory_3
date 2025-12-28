#pragma once
#include <string>
#include <fstream>
#include <limits>

#include "json.hpp"

using namespace std;
using json = nlohmann::json;

template <typename DS>
void runInteractive(const string& name) {
    DS ds;

    ifstream in(name + ".json");
    if (in) {
        json j;
        in >> j;
        ds.from_json(j);
        cout << "< Loaded from " << name << ".json\n";
    }

    while (true) {
        cout << "> ";
        string cmd;
        cin >> cmd;

        if (cmd == "add") {
            int x;
            cin >> x;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
            ds.push_back(x);
            cout << "< OK\n";
        }
        else if (cmd == "remove") {
            int x;
            cin >> x;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
            ds.remove(x);
            cout << "< OK\n";
        }
        else if (cmd == "find") {
            int x;
            cin >> x;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
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
            json j;
            ds.to_json(j);
            ofstream out(name + ".json");
            if (out) {
                out << j.dump(4);
                cout << "< Autosaved to " << name << ".json\n";
            }
            break;
        }
        else {
            cout << "< Unknown command\n";
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cin.clear();
        }
    }
}

template <typename HashMap>
void runInteractiveHash(const string& name) {
    HashMap map;
    
    ifstream in(name + ".json");
    if (in) {
        json j;
        in >> j;
        map.from_json(j);
        cout << "< Loaded from " << name << ".json\n";
    }

    while (true) {
        cout << "> ";
        string cmd;
        cin >> cmd;

        if (cmd == "add") {
            int k, v;
            cin >> k >> v;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
            map.put(k, v);
            cout << "< OK\n";
        }
        else if (cmd == "remove") {
            int k;
            cin >> k;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
            map.remove(k);
            cout << "< OK\n";
        }
        else if (cmd == "find") {
            int k;
            cin >> k;
            if (cin.fail()) {
                cout << "< Invalid input\n";
                cin.clear();
                cin.ignore(numeric_limits<streamsize>::max(), '\n');
                continue;
            }
            cout << (map.contains(k) ? "< Exists\n" : "< Not found\n");
        }
        else if (cmd == "print") {
            map.display();
        }
        else if (cmd == "exit") {
            json j;
            map.to_json(j);
            ofstream out(name + ".json");
            if (out) {
                out << j.dump(4);
                cout << "< Autosaved to " << name << ".json\n";
            }
            break;
        }
        else {
            cout << "< Unknown command\n";
            cin.ignore(numeric_limits<streamsize>::max(), '\n');
            cin.clear();
        }
    }
}
