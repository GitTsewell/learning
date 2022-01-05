##值类型
+ bool
+ int/uint(8~256)
+ 定点数 fixed/ufixed (目前并没有完全支持)
+ 地址 address(20个字节)  address payable (可以将以太坊发送到该地址), address 转换到 address payable 使用payable(address)
    + 地址成员 balance transfer . address.balance 该地址余额 , address.transfer(10) 向该地址转账(以 wei为单位)
    + transfer: 如果address是一个合约,transfer调用时(fallback function将一起调用),如果transfer失败.以太币转移将被恢复,合约异常停止
    + send : send是低级的transfer,如果执行失败,当前合约不会因异常而停止,send会返回false
    + call,delegatecall, staticcall,直接调用合约address的function
      + address.call(abi.encodewithsignature("register(string)", "myName"))
      + 也可以调整gas费 address.call{gas : 10000}(abi.encodewithsignature("register(string)", "myName"))
      + 也可以调整以太币值 address(nameReg).call{value: 1 ether}(abi.encodeWithSignature("register(string)", "MyName"));
+ bytes
+ string
+ unicode 文字 unicode可以包含任何有效的 UTF-8 序列 example : string memory a = unicode"Hello 😃";
+ 十六进制文字 hex
+ 枚举 enum . example : enum ActionChoices { GoLeft, GoRight, GoStraight, SitStill }  , ActionChoices.GoLeft
+ 函数类型
  + 外部函数 public external
  + 内部函数 internal 只能在当前合约内部调用,包括内部函数和继承函数
  + 函数转换
    + pure函数可以转换为view和non-payable函数
    + view函数可以转换为non-payable函数
    + payable函数可以转换为non-payable函数
  + 函数成员
    + .address 返回函数合约的地址
    + .selector 返回ABI函数选择器

##引用类型
+ memory 生命周期仅限于外部函数调用
+ storage 生命周期仅限于合约的生命周期
+ calldata 不可修改,非持久性区域,用于存储函数参数,其行为主要类似于内存
+ array 可以固定大小,也可以动态大小 方法.push用于追加末尾元素
+ struct
+ map

##单位
###以太单位
+ wei : 1 wei == 1
+ gwei : 1gwei == 1e9
+ ether : 1ether == 1e18

###时间单位
```
seconds,minutes,hours,days,weeks
```

##特殊变量和函数
###区块和交易属性
+ blockhash(uint blockNumber) returns (bytes32): 给定块的哈希值，当blocknumber是 256 个最近的块之一时；否则返回零
+ block.basefee( uint): 当前区块的基本费用（EIP-3198和EIP-1559）
+ block.chainid( uint): 当前链 id
+ block.coinbase( ): 当前区块矿工地址address payable
+ block.difficulty( uint): 当前区块难度
+ block.gaslimit( uint): 当前区块的gaslimit
+ block.number( uint): 当前区块号
+ block.timestamp( uint): 当前区块时间戳，自 Unix 纪元以来的秒数
+ gasleft() returns (uint256): 剩余气体
+ msg.data( ): 完整的调用数据bytes calldata
+ msg.sender( address): 消息的发送者（当前调用）
+ msg.sig( bytes4): calldata 的前四个字节（即函数标识符）
+ msg.value( uint): 随消息发送的wei数量
+ tx.gasprice( uint): 交易的gas价格
+ tx.origin( address): 交易的发送者（完整的调用链）

###ABI编码和解码函数
+ abi.decode(bytes memory encodedData, (...)) returns (...): ABI 解码给定的数据，而类型在括号中作为第二个参数给出
+ abi.encode(...) returns (bytes memory): ABI 编码给定的参数
+ abi.encodePacked(...) returns (bytes memory): 执行给定参数的压缩编码。请注意，打包编码可能不明确！
+ abi.encodeWithSelector(bytes4 selector, ...) returns (bytes memory): ABI 从第二个开始对给定的参数进行编码，并在给定的四字节选择器之前
+ abi.encodeWithSignature(string memory signature, ...) returns (bytes memory)： 相当于 abi.encodeWithSelector(bytes4(keccak256(bytes(signature))), ...)`

###字节成员
```
bytes.concat(...) returns (bytes memory): 将可变数量的字节和 bytes1, ..., bytes32 参数连接到一个字节数组
```
###错误处理
+ assert(bool condition) 如果条件不满足，则会导致 Panic 错误，从而恢复状态更改 - 用于内部错误。
+ require(bool condition) 如果条件不满足，则恢复 - 用于输入或外部组件中的错误
+ require(bool condition, string memory message) 如果条件不满足，则恢复 - 用于输入或外部组件中的错误。还提供错误信息。
+ revert() 中止执行并恢复状态更改
+ revert(string memory reason) 中止执行并恢复状态更改，提供解释性字符串

###数学和密码函数
+ addmod(uint x, uint y, uint k) returns (uint) ; X加Y 然后对K 取余
+ mulmod(uint x, uint y, uint k) returns (uint) ; X乘Y 然后对K 取余
+ keccak256(bytes memory) returns (bytes32) 计算输入的 Keccak-256 哈希
+ sha256(bytes memory) returns (bytes32) 计算输入的 SHA-256 哈希值
+ ripemd160(bytes memory) returns (bytes20) 计算输入的 RIPEMD-160 哈希

###地址类型的成员
+ address.balance( uint256) 地址的余额 单位wei
+ address.code( )bytes memory 地址处的代码（可以为空）
+ address.codehash( bytes32) 地址的代码哈希
+ address payable.transfer(uint256 amount)发送给定数量的 Wei 到Address，失败时恢复，转发 2300 gas 津贴，不可调整
+ address payable.send(uint256 amount) returns (bool) 发送给定数量的 Wei 到Address，false失败返回，转发 2300 gas 津贴，不可调整
+ address.call(bytes memory) returns (bool, bytes memory) CALL使用给定的有效载荷发出低级别，返回成功条件和返回数据，转发所有可用的气体，可调整

####note
```
在 0.5.0 版本之前，Solidity 允许合约实例访问地址成员，例如this.balance。现在禁止这样做，必须进行显式转换为地址：address(this).balance。
```

###合同相关
+ this ; 当前合约
+ selfdestruct(address payable recipient) 销毁当前合约，将其资金发送到给定地址 并结束执行。请注意，它selfdestruct有一些继承自 EVM 的特性：
接收合约的接收函数没有执行。 合约仅在交易结束时才真正被销毁，reverts 可能会“撤消”销毁

###类型信息
+ type(C).name 合同名称。
+ type(C).creationCode 包含合约创建字节码的内存字节数组。这可用于内联汇编以构建自定义创建例程，尤其是通过使用create2操作码。此属性不能在合约本身或任何派生合约中访问。它导致字节码包含在调用站点的字节码中，因此不可能进行这样的循环引用
+ type(C).runtimeCode 包含合约运行时字节码的内存字节数组。这是通常由 的构造函数部署的代码C。如果C有一个使用内联汇编的构造函数，这可能与实际部署的字节码不同。另请注意，库在部署时修改其运行时字节码以防止常规调用。与 相同的限制.creationCode也适用于此属性。
+ type(I).interfaceId： 甲bytes4包含值EIP-165 给定的接口的接口标识符I。该标识符被定义为XOR接口本身内定义的所有函数选择器的标识符——不包括所有继承的函数。
+ type(T).min type 可表示的最小值T
+ type(T).max type 可表示的最大值T

##错误处理 Assert、Require、Revert 和 Exceptions
###Assert
该assert函数会创建类型为 的错误Panic(uint256),断言应该只用于测试内部错误和检查不变量。正常运行的代码永远不应该造成恐慌，即使是在无效的外部输入上也不应该.

以下情况会产生 Panic 异常。与错误数据一起提供的错误代码表示恐慌的类型
1. 0x00：用于通用编译器插入的恐慌
2. 0x01：如果您assert使用计算结果为 false 的参数调用
3. 0x11：如果算术运算导致块外下溢或上溢。unchecked { ... }
4. 0x12: 如果您除以或取模为零（例如或）。5 / 023 % 0
5. 0x21：如果将太大或为负的值转换为枚举类型。
6. 0x22：如果您访问的存储字节数组编码不正确。
7. 0x31：如果你调用.pop()一个空数组。
8. 0x32：如果您bytesN在越界或负索引（即x[i]where或）处访问数组或数组切片。i >= x.lengthi < 0
9. 0x41：如果分配太多内存或创建的数组太大。
10. 0x51：如果调用内部函数类型的零初始化变量。

### require
该require函数要么创建一个没有任何数据的错误，要么创建一个类型为 的错误Error(string)。它应该用于确保在执行时间之前无法检测到的有效条件。这包括对外部合同调用的输入或返回值的条件。

一个Error(string)例外（或不带数据的异常）由在下列情况下，编译器生成的：
1. 调用require(x)wherex评估为false.
2. 如果您使用revert()或revert("description")。
3. 如果您执行针对不包含代码的合同的外部函数调用。
4. 如果您的合约通过没有payable修饰符的公共函数（包括构造函数和回退函数）接收 Ether 。
5. 如果您的合约通过公共 getter 函数接收 Ether。

对于以下情况，将转发来自外部调用（如果提供）的错误数据。这意味着它可能导致错误或恐慌（或其他任何给出的）：
1. 如果一个.transfer()失败。
2. 如果您通过消息调用调用一个函数，但它没有正确完成（即，它耗尽了 gas，没有匹配的函数，或者本身抛出异常），除非使用了低级操作 call、send、delegatecall、callcode或staticcall 。低级操作从不抛出异常，而是通过返回指示失败false。
3. 如果您使用new关键字创建合同，但合同创建未正确完成。

### revert
if(msg.sender != owner) { revert(); }  等同于 require(msg.sender == owner);





## contract
### 创建合约
1. 可以通过eth transaction 上链创建合约
2. 也可以在solidity内部创建合约,通过constructor调用创建
3. 通过程序创建的另一种方式,还可以通过web3.js的web3.eth.Contract函数创建

### 可见性
1. external : 外部函数是合约接口的一部分，这意味着它们可以从其他合约和通过交易调用。外部函数f不能在内部调用（即f()不起作用，但this.f()有效）
2. public : 公共函数是合约接口的一部分，可以在内部或通过消息调用。对于公共状态变量，会生成一个自动 getter 函数（见下文）。
3. internal : 这些函数和状态变量只能在内部访问（即从当前合约或从它继承的合约中），而不能使用this. 这是状态变量的默认可见性级别
4. private : 私有函数和状态变量仅对定义它们的合约可见，在派生合约中不可见。

### getter 访问器
对于所有public的状态变量，Solidity语言编译器，提供了自动为状态变量生成对应的getter（访问器）的特性

### modifier 函数修饰符
修饰符可用于以声明方式更改函数的行为。例如，您可以使用修饰符在执行函数之前自动检查条件。

修饰符是合约的可继承属性，可以被派生合约覆盖，但前提是它们被标记为virtual

### 常量和不可变状态变量 
状态变量可以声明为constant或immutable。在这两种情况下，在合约构建后无法修改变量。对于constant变量，其值必须在编译时固定，而对于immutable，它仍然可以在构造时赋值。

与常规状态变量相比，常量和不可变变量的 gas 成本要低得多。对于常量变量，分配给它的表达式被复制到它被访问的所有地方，并且每次都重新计算。这允许局部优化。不可变变量在构造时被评估一次，它们的值被复制到代码中访问它们的所有位置。对于这些值，保留 32 个字节，即使它们适合较少的字节。因此，常量值有时比不可变值便宜。

并非所有常量和不可变类型都在此时实现。唯一支持的类型是 字符串（仅用于常量）和值类型。

#### constant
对于constant变量，该值在编译时必须是常量，并且必须在声明变量的地方分配。任何访问存储、区块链数据（例如block.timestamp，address(this).balance或 block.number）或执行数据（msg.value或gasleft()）或调用外部合约的表达式都是不允许的。允许可能对内存分配产生副作用的表达式，但不允许对其他内存对象产生副作用的表达式。内置的功能 keccak256，sha256，ripemd160，ecrecover，addmod和mulmod 被允许的（尽管，与例外keccak256，他们调用外部合同）。

#### immutable
声明为的变量immutable比声明为的constant变量限制要少一些：可以在合约的构造函数中或在声明点为不可变变量分配任意值。它们只能分配一次，从那时起，即使在构建期间也可以读取。

编译器生成的合约创建代码将在合约返回之前修改合约的运行时代码，方法是将所有对不可变项的引用替换为分配给它们的值。如果您将编译器生成的运行时代码与实际存储在区块链中的代码进行比较，这一点很重要

### functions
函数可以在合约内部和外部定义。

合约之外的函数，也称为“自由函数”，总是具有隐式internal 可见性。它们的代码包含在调用它们的所有合约中，类似于内部库函数。

#### 函数状态可变性
##### view 查看
声明函数，view 在这种情况下，它们承诺不修改状态。

以下事件被认为是修改状态
1. Writing to state variables.
2. Emitting events.
3. Creating other contracts.
4. Using selfdestruct.
5. Sending Ether via calls.
6. Calling any function not marked view or pure.
7. Using low-level calls.
8. 使用包含特定操作码的内联汇编。

##### pure 纯函数
可以声明函数，pure在这种情况下，它们承诺不读取或修改状态。特别是，应该可以pure在编译时评估一个函数，只给出它的输入 和msg.data，但不知道当前区块链状态。这意味着从immutable变量中读取可以是非纯操作


以下事件被任务是从状态中读取
1. Reading from state variables.
2. Accessing address(this).balance or <address>.balance
3. Accessing any of the members of block, tx, msg (with the exception of msg.sig and msg.data).
4. Calling any function not marked pure.
5. 使用包含特定操作码的内联汇编

#### Special Functions
##### Receive
一个合约最多可以有一个receive()函数，使用 receive() external payable { ... } （没有 function 关键字）声明。此函数不能有参数，不能返回任何内容，并且必须具有 external 和 payable 。它可以是虚拟的，可以 override 并且可以具有 modifiers

接收函数在调用具有空 calldata 的合约时执行。这是在普通 Ether 传输（例如 via.send()或.transfer()）上执行的函数。如果不存在这样的函数，但存在可支付的回退函数 ，则回退函数将在普通 Ether 传输中被调用。如果不存在接收以太币和应付回退功能，则合约无法通过常规交易接收以太币并引发异常。

##### fallback
一个合约最多可以有一个回退函数，使用 fallback () external [payable] 或 fallback (bytes calldata _input) external [payable] return (bytes memory _output) 声明（两者都没有 function 关键字）。此功能必须具有外部可见性。回退函数可以是虚拟的，可以覆盖并且可以具有修饰符

如果没有其他函数与给定的函数签名匹配，或者根本没有提供数据并且没有接收以太函数，则在调用合约时执行回退函数。回退函数总是接收数据，但为了也接收以太，它必须被标记为payable。

##### overload
一个合约可以有多个同名但参数类型不同的函数。这个过程称为“重载”，也适用于继承的函数
```
pragma solidity >=0.4.16 <0.9.0;

contract A {
    function f(uint _in) public pure returns (uint out) {
        out = _in;
    }

    function f(uint _in, bool _really) public pure returns (uint out) {
        if (_really)
            out = _in;
    }
}
```






### events
Solidity 事件在 EVM 的日志记录功能之上提供了一个抽象。应用程序可以通过以太坊客户端的 RPC 接口订阅和监听这些事件

事件是合约的可继承成员。当你调用它们时，它们会导致参数存储在交易日志中——区块链中的一种特殊数据结构。这些日志与合约的地址相关联，被合并到区块链中，并且只要一个区块可访问就一直存在（直到现在永远存在，但这可能会随着 Serenity 改变）。无法从合约内部访问日志及其事件数据（甚至不能从创建它们的合约中访问）。

### Function Overriding
如果基函数被标记为 ，则可以通过继承契约来覆盖它们以更改它们的行为virtual。然后覆盖函数必须使用override函数头中的关键字。覆盖函数只能将覆盖函数的可见性从external更改为public。可变性可能会按照以下顺序更改为更严格的： nonpayable可以被view和覆盖pure。view可以被覆盖pure。 payable是一个例外，不能更改为任何其他可变性。

```
// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract Base
{
    function foo() virtual external view {}
}

contract Middle is Base {}

contract Inherited is Middle
{
    function foo() override public pure {}
}
```

### Constructors
构造函数是用constructor关键字声明的可选函数，它在合约创建时执行，您可以在其中运行合约初始化代码。

构造函数运行后，合约的最终代码被部署到区块链。代码的部署成本与代码长度成线性关系。此代码包括属于公共接口的所有函数以及可通过函数调用从那里访问的所有函数。它不包括构造函数代码或仅从构造函数调用的内部函数。


