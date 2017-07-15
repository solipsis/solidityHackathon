pragma solidity ^0.4.4;
contract Transfer {

    uint public threshold;
    address public owner;
    bool public funded;

    function Transfer(uint _threshold) {
        owner = msg.sender;
        threshold = _threshold;
        funded = false;
    }

    function payOut(address addr) payable {
        if (msg.sender != owner) {
            addr.send(this.balance);
        }
    }

    function() payable {
        if (this.balance >= threshold) {
           funded = true;
        }
    }
}

contract Spawn {

    struct Issue {
        uint threshold;
        string id;
    }

    mapping(address => Issue) public issues;
    address[] public addressLUT; 

    function createIssue(string id, uint threshold) returns (address) {
        address subAddress = new Transfer(threshold);

        issues[subAddress].id = id;
        issues[subAddress].threshold = threshold;
        addressLUT.push(subAddress);
        return subAddress;
    }

    function size() constant returns (uint) {
        return addressLUT.length;
    }

    

    function() payable {
        
    }
}