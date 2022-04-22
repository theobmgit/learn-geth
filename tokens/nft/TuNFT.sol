// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../../packages/openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "../../packages/openzeppelin/contracts/access/Ownable.sol";
import "../../packages/openzeppelin/contracts/utils/Counters.sol";

contract TuNFT is ERC721URIStorage, Ownable {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor() ERC721("TuNFT", "NFT") {}

    function mintNFT(address recipient, string memory tokenURI)
    public onlyOwner
    returns (uint256)
    {
        _tokenIds.increment();

        uint256 newItemId = _tokenIds.current();
        _mint(recipient, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }
}
