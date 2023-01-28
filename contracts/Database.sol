//SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;


struct Car {
    string category;
    uint32 cost;
    uint32 ageCount;
}

struct Document{
    string number;
    uint64 validateTime;
    string category;
}

struct Driver {
    string fio;
    uint8 exp;
    uint32 dtpCount;
    uint32 fineCount;
}

struct LoginData {
    address addr;
    string password;
    string role;
}


contract Database {

    address payable public bank;

    address payable public insurance;

    mapping(string => LoginData) public logins;

    mapping(string => address) public driverDocNumbers; 

    mapping(address => Document) public driverDocuments;
    
    mapping(address => Driver) public drivers;

    mapping(address => Car) public driverCars;

    mapping(address => uint[]) public driverFines;

    mapping(address => uint) public driverInsurances;

    uint public insuranceLoan; 

    

    constructor(address _bank,address _insurance) {
        bank = payable(_bank);
        LoginData memory bankLoginData = LoginData({addr:bank,password:"bank",role:"bank"});
        logins["bank"] = bankLoginData;
        insurance = payable(_insurance);
        LoginData memory insLoginData = LoginData({addr:_insurance,password:"insurance",role:"insurance"});
        logins["insurance"] = insLoginData;
    }

    function addDriver(
        string memory fio,
        uint8 exp,
        uint32 dtpCount,
        uint32 fineCount,
        string memory login,
        string memory password
    ) public {
        LoginData memory newLoginData = LoginData({addr:msg.sender,password:password,role:"driver"});
        Driver memory newDriver = Driver({fio:fio,exp:exp,dtpCount:dtpCount,fineCount:fineCount});
        drivers[msg.sender] = newDriver;
        logins[login] = newLoginData;
        for (uint i = 0; i < fineCount; i++){
            driverFines[msg.sender].push(block.timestamp);
        }
    }

    function addPoliceMan(
        string memory fio,
        uint8 exp,
        uint32 dtpCount,
        uint32 fineCount,
        string memory login,
        string memory password
    ) public {
        LoginData memory newLoginData = LoginData({addr:msg.sender,password:password,role:"policeman"});
        Driver memory newDriver = Driver({fio:fio,exp:exp,dtpCount:dtpCount,fineCount:fineCount});
        drivers[msg.sender] = newDriver;
        logins[login] = newLoginData;
        for (uint i = 0; i < fineCount; i++){
            driverFines[msg.sender].push(block.timestamp);
        }
    }

    function addDriverCar(string memory category, uint32 cost,uint32 ageCount) public {
        Car memory newCar = Car({category:category,cost:cost,ageCount:ageCount});
        driverCars[msg.sender] = newCar;
    }

    function addDriverDocument(string memory number,uint64 validateTime,string memory category) public {
        Document memory newDocument = Document({number:number,validateTime:validateTime,category:category});
        driverDocNumbers[number] = msg.sender;
        driverDocuments[msg.sender] = newDocument;
    }

    function payFine() public payable{
        uint[] memory fines = driverFines[msg.sender];
        int checkTime = int(block.timestamp) - int(fines[fines.length - 1]);
        if (checkTime > 25){
            bank.transfer(10 ether);
            driverFines[msg.sender].pop();
            drivers[msg.sender].fineCount -= 1;
        }else {
            bank.transfer(5 ether);
            payable(msg.sender).transfer(5 ether);
            driverFines[msg.sender].pop();
            drivers[msg.sender].fineCount -= 1;
        }
    }

    function addFine(string memory number) public {
        address drAddr = driverDocNumbers[number];
        require(drAddr != address(0));
        driverFines[drAddr].push(block.timestamp);
        drivers[drAddr].fineCount += 1;
    }

    function buyInsurance(uint amount) public payable {
        require(payable(msg.sender).balance >= amount,"low balance");
        require(driverInsurances[msg.sender] == 0,"You have insurance");
        if (insuranceLoan > amount){
            bank.transfer(amount);
            driverInsurances[msg.sender] = amount;
            insuranceLoan -= amount;
        } else if (insuranceLoan < amount){
            uint diff = amount - insuranceLoan;
            bank.transfer(insuranceLoan);
            insurance.transfer(diff);
            driverInsurances[msg.sender] = amount;
            insuranceLoan = 0;
        } else if (insuranceLoan == amount) {
            bank.transfer(amount);
            driverInsurances[msg.sender] = amount;
            insuranceLoan = 0;
        }
    }

    function renewInsurance() public {
        Document memory dc = driverDocuments[msg.sender];
        require((dc.validateTime - block.timestamp) > 30 days, "wrong time");
        require(driverFines[msg.sender].length == 0, "You have fines");
        driverDocuments[msg.sender].validateTime += 10 * 365 days ;
    }  

    function addDTP(string memory number) public {
        address drAddr = driverDocNumbers[number];
        require(drAddr != address(0));
        drivers[drAddr].dtpCount += 1;
    }

    function incLoan(uint amount)public{
        insuranceLoan += amount;
    }
    function decLoan(uint amount)public{
        insuranceLoan -= amount;
    }
    function setLoan(uint amount)public{
        insuranceLoan = amount;
    }




}