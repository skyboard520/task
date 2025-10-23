// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyImageNFT is ERC721, Ownable {
    uint256 private _tokenIdCounter = 1;

    // 存储每个 token 的 URI
    mapping(uint256 => string) private _tokenURIs;
    //构造函数：初始化NFT名称和符号
    constructor(string memory name_, string memory symbol_) ERC721(name_, symbol_) Ownable(msg.sender){
    }

    // 铸造NFT函数（仅所有者可调用）
    //tokenURL:ipfs://bafkreigx3uqhvoxkejqdoyenohbk5neubuagmvipkaq6qwp3nrcml2o3mq
    function mintNFT(address recipient, string memory tokenURI) public onlyOwner returns (uint256) {
        uint256 newTokenId = _tokenIdCounter;
        _safeMint(recipient, newTokenId);
        _tokenURIs[newTokenId] = tokenURI;
        _tokenIdCounter++;
        return newTokenId;
    }

    // 重写tokenURI函数（OpenZeppelin v4.x及以上需要显式实现）
    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        _requireOwned(tokenId); // 检查 token 是否存在
        return _tokenURIs[tokenId];
    }
}