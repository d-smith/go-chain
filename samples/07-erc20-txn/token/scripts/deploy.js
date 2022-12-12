const hre = require("hardhat");

async function main() {

  const TestToken = await hre.ethers.getContractFactory("TestToken");
  const token = await TestToken.deploy();
  await token.deployed();
  console.log( ` deployed to ${token.address}` );

}

main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
