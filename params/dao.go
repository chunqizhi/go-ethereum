// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"

	"github.com/chunqizhi/go-ethereum/common"
)

// DAOForkBlockExtra is the block header extra-data field to set for the DAO fork
// point and a number of consecutive blocks to allow fast/light syncers to correctly
// pick the side they want  ("dao-hard-fork").
var DAOForkBlockExtra = common.FromHex("Gs64616f2d686172642d666f726b")

// DAOForkExtraRange is the number of consecutive blocks from the DAO fork point
// to override the extra-data in to prevent no-fork attacks.
var DAOForkExtraRange = big.NewInt(10)

// DAORefundContract is the address of the refund contract to send DAO balances to.
var DAORefundContract = common.HexToAddress("Gsbf4ed7b27f1d666546e30d74d50d173d20bca754")

// DAODrainList is the list of accounts whose full balances will be moved into a
// refund contract at the beginning of the dao-fork block.
func DAODrainList() []common.Address {
	return []common.Address{
		common.HexToAddress("Gsd4fe7bc31cedb7bfb8a345f31e668033056b2728"),
		common.HexToAddress("Gsb3fb0e5aba0e20e5c49d252dfd30e102b171a425"),
		common.HexToAddress("Gs2c19c7f9ae8b751e37aeb2d93a699722395ae18f"),
		common.HexToAddress("Gsecd135fa4f61a655311e86238c92adcd779555d2"),
		common.HexToAddress("Gs1975bd06d486162d5dc297798dfc41edd5d160a7"),
		common.HexToAddress("Gsa3acf3a1e16b1d7c315e23510fdd7847b48234f6"),
		common.HexToAddress("Gs319f70bab6845585f412ec7724b744fec6095c85"),
		common.HexToAddress("Gs06706dd3f2c9abf0a21ddcc6941d9b86f0596936"),
		common.HexToAddress("Gs5c8536898fbb74fc7445814902fd08422eac56d0"),
		common.HexToAddress("Gs6966ab0d485353095148a2155858910e0965b6f9"),
		common.HexToAddress("Gs779543a0491a837ca36ce8c635d6154e3c4911a6"),
		common.HexToAddress("Gs2a5ed960395e2a49b1c758cef4aa15213cfd874c"),
		common.HexToAddress("Gs5c6e67ccd5849c0d29219c4f95f1a7a93b3f5dc5"),
		common.HexToAddress("Gs9c50426be05db97f5d64fc54bf89eff947f0a321"),
		common.HexToAddress("Gs200450f06520bdd6c527622a273333384d870efb"),
		common.HexToAddress("Gsbe8539bfe837b67d1282b2b1d61c3f723966f049"),
		common.HexToAddress("Gs6b0c4d41ba9ab8d8cfb5d379c69a612f2ced8ecb"),
		common.HexToAddress("Gsf1385fb24aad0cd7432824085e42aff90886fef5"),
		common.HexToAddress("Gsd1ac8b1ef1b69ff51d1d401a476e7e612414f091"),
		common.HexToAddress("Gs8163e7fb499e90f8544ea62bbf80d21cd26d9efd"),
		common.HexToAddress("Gs51e0ddd9998364a2eb38588679f0d2c42653e4a6"),
		common.HexToAddress("Gs627a0a960c079c21c34f7612d5d230e01b4ad4c7"),
		common.HexToAddress("Gsf0b1aa0eb660754448a7937c022e30aa692fe0c5"),
		common.HexToAddress("Gs24c4d950dfd4dd1902bbed3508144a54542bba94"),
		common.HexToAddress("Gs9f27daea7aca0aa0446220b98d028715e3bc803d"),
		common.HexToAddress("Gsa5dc5acd6a7968a4554d89d65e59b7fd3bff0f90"),
		common.HexToAddress("Gsd9aef3a1e38a39c16b31d1ace71bca8ef58d315b"),
		common.HexToAddress("Gs63ed5a272de2f6d968408b4acb9024f4cc208ebf"),
		common.HexToAddress("Gs6f6704e5a10332af6672e50b3d9754dc460dfa4d"),
		common.HexToAddress("Gs77ca7b50b6cd7e2f3fa008e24ab793fd56cb15f6"),
		common.HexToAddress("Gs492ea3bb0f3315521c31f273e565b868fc090f17"),
		common.HexToAddress("Gs0ff30d6de14a8224aa97b78aea5388d1c51c1f00"),
		common.HexToAddress("Gs9ea779f907f0b315b364b0cfc39a0fde5b02a416"),
		common.HexToAddress("Gsceaeb481747ca6c540a000c1f3641f8cef161fa7"),
		common.HexToAddress("Gscc34673c6c40e791051898567a1222daf90be287"),
		common.HexToAddress("Gs579a80d909f346fbfb1189493f521d7f48d52238"),
		common.HexToAddress("Gse308bd1ac5fda103967359b2712dd89deffb7973"),
		common.HexToAddress("Gs4cb31628079fb14e4bc3cd5e30c2f7489b00960c"),
		common.HexToAddress("Gsac1ecab32727358dba8962a0f3b261731aad9723"),
		common.HexToAddress("Gs4fd6ace747f06ece9c49699c7cabc62d02211f75"),
		common.HexToAddress("Gs440c59b325d2997a134c2c7c60a8c61611212bad"),
		common.HexToAddress("Gs4486a3d68fac6967006d7a517b889fd3f98c102b"),
		common.HexToAddress("Gs9c15b54878ba618f494b38f0ae7443db6af648ba"),
		common.HexToAddress("Gs27b137a85656544b1ccb5a0f2e561a5703c6a68f"),
		common.HexToAddress("Gs21c7fdb9ed8d291d79ffd82eb2c4356ec0d81241"),
		common.HexToAddress("Gs23b75c2f6791eef49c69684db4c6c1f93bf49a50"),
		common.HexToAddress("Gs1ca6abd14d30affe533b24d7a21bff4c2d5e1f3b"),
		common.HexToAddress("Gsb9637156d330c0d605a791f1c31ba5890582fe1c"),
		common.HexToAddress("Gs6131c42fa982e56929107413a9d526fd99405560"),
		common.HexToAddress("Gs1591fc0f688c81fbeb17f5426a162a7024d430c2"),
		common.HexToAddress("Gs542a9515200d14b68e934e9830d91645a980dd7a"),
		common.HexToAddress("Gsc4bbd073882dd2add2424cf47d35213405b01324"),
		common.HexToAddress("Gs782495b7b3355efb2833d56ecb34dc22ad7dfcc4"),
		common.HexToAddress("Gs58b95c9a9d5d26825e70a82b6adb139d3fd829eb"),
		common.HexToAddress("Gs3ba4d81db016dc2890c81f3acec2454bff5aada5"),
		common.HexToAddress("Gsb52042c8ca3f8aa246fa79c3feaa3d959347c0ab"),
		common.HexToAddress("Gse4ae1efdfc53b73893af49113d8694a057b9c0d1"),
		common.HexToAddress("Gs3c02a7bc0391e86d91b7d144e61c2c01a25a79c5"),
		common.HexToAddress("Gs0737a6b837f97f46ebade41b9bc3e1c509c85c53"),
		common.HexToAddress("Gs97f43a37f595ab5dd318fb46e7a155eae057317a"),
		common.HexToAddress("Gs52c5317c848ba20c7504cb2c8052abd1fde29d03"),
		common.HexToAddress("Gs4863226780fe7c0356454236d3b1c8792785748d"),
		common.HexToAddress("Gs5d2b2e6fcbe3b11d26b525e085ff818dae332479"),
		common.HexToAddress("Gs5f9f3392e9f62f63b8eac0beb55541fc8627f42c"),
		common.HexToAddress("Gs057b56736d32b86616a10f619859c6cd6f59092a"),
		common.HexToAddress("Gs9aa008f65de0b923a2a4f02012ad034a5e2e2192"),
		common.HexToAddress("Gs304a554a310c7e546dfe434669c62820b7d83490"),
		common.HexToAddress("Gs914d1b8b43e92723e64fd0a06f5bdb8dd9b10c79"),
		common.HexToAddress("Gs4deb0033bb26bc534b197e61d19e0733e5679784"),
		common.HexToAddress("Gs07f5c1e1bc2c93e0402f23341973a0e043f7bf8a"),
		common.HexToAddress("Gs35a051a0010aba705c9008d7a7eff6fb88f6ea7b"),
		common.HexToAddress("Gs4fa802324e929786dbda3b8820dc7834e9134a2a"),
		common.HexToAddress("Gs9da397b9e80755301a3b32173283a91c0ef6c87e"),
		common.HexToAddress("Gs8d9edb3054ce5c5774a420ac37ebae0ac02343c6"),
		common.HexToAddress("Gs0101f3be8ebb4bbd39a2e3b9a3639d4259832fd9"),
		common.HexToAddress("Gs5dc28b15dffed94048d73806ce4b7a4612a1d48f"),
		common.HexToAddress("Gsbcf899e6c7d9d5a215ab1e3444c86806fa854c76"),
		common.HexToAddress("Gs12e626b0eebfe86a56d633b9864e389b45dcb260"),
		common.HexToAddress("Gsa2f1ccba9395d7fcb155bba8bc92db9bafaeade7"),
		common.HexToAddress("Gsec8e57756626fdc07c63ad2eafbd28d08e7b0ca5"),
		common.HexToAddress("Gsd164b088bd9108b60d0ca3751da4bceb207b0782"),
		common.HexToAddress("Gs6231b6d0d5e77fe001c2a460bd9584fee60d409b"),
		common.HexToAddress("Gs1cba23d343a983e9b5cfd19496b9a9701ada385f"),
		common.HexToAddress("Gsa82f360a8d3455c5c41366975bde739c37bfeb8a"),
		common.HexToAddress("Gs9fcd2deaff372a39cc679d5c5e4de7bafb0b1339"),
		common.HexToAddress("Gs005f5cee7a43331d5a3d3eec71305925a62f34b6"),
		common.HexToAddress("Gs0e0da70933f4c7849fc0d203f5d1d43b9ae4532d"),
		common.HexToAddress("Gsd131637d5275fd1a68a3200f4ad25c71a2a9522e"),
		common.HexToAddress("Gsbc07118b9ac290e4622f5e77a0853539789effbe"),
		common.HexToAddress("Gs47e7aa56d6bdf3f36be34619660de61275420af8"),
		common.HexToAddress("Gsacd87e28b0c9d1254e868b81cba4cc20d9a32225"),
		common.HexToAddress("Gsadf80daec7ba8dcf15392f1ac611fff65d94f880"),
		common.HexToAddress("Gs5524c55fb03cf21f549444ccbecb664d0acad706"),
		common.HexToAddress("Gs40b803a9abce16f50f36a77ba41180eb90023925"),
		common.HexToAddress("Gsfe24cdd8648121a43a7c86d289be4dd2951ed49f"),
		common.HexToAddress("Gs17802f43a0137c506ba92291391a8a8f207f487d"),
		common.HexToAddress("Gs253488078a4edf4d6f42f113d1e62836a942cf1a"),
		common.HexToAddress("Gs86af3e9626fce1957c82e88cbf04ddf3a2ed7915"),
		common.HexToAddress("Gsb136707642a4ea12fb4bae820f03d2562ebff487"),
		common.HexToAddress("Gsdbe9b615a3ae8709af8b93336ce9b477e4ac0940"),
		common.HexToAddress("Gsf14c14075d6c4ed84b86798af0956deef67365b5"),
		common.HexToAddress("Gsca544e5c4687d109611d0f8f928b53a25af72448"),
		common.HexToAddress("Gsaeeb8ff27288bdabc0fa5ebb731b6f409507516c"),
		common.HexToAddress("Gscbb9d3703e651b0d496cdefb8b92c25aeb2171f7"),
		common.HexToAddress("Gs6d87578288b6cb5549d5076a207456a1f6a63dc0"),
		common.HexToAddress("Gsb2c6f0dfbb716ac562e2d85d6cb2f8d5ee87603e"),
		common.HexToAddress("Gsaccc230e8a6e5be9160b8cdf2864dd2a001c28b6"),
		common.HexToAddress("Gs2b3455ec7fedf16e646268bf88846bd7a2319bb2"),
		common.HexToAddress("Gs4613f3bca5c44ea06337a9e439fbc6d42e501d0a"),
		common.HexToAddress("Gsd343b217de44030afaa275f54d31a9317c7f441e"),
		common.HexToAddress("Gs84ef4b2357079cd7a7c69fd7a37cd0609a679106"),
		common.HexToAddress("Gsda2fef9e4a3230988ff17df2165440f37e8b1708"),
		common.HexToAddress("Gsf4c64518ea10f995918a454158c6b61407ea345c"),
		common.HexToAddress("Gs7602b46df5390e432ef1c307d4f2c9ff6d65cc97"),
		common.HexToAddress("Gsbb9bc244d798123fde783fcc1c72d3bb8c189413"),
		common.HexToAddress("Gs807640a13483f8ac783c557fcdf27be11ea4ac7a"),
	}
}
