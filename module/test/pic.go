package main

import (
	"encoding/hex"
	"fyne.io/fyne/v2"
)
var Logo,_ = hex.DecodeString("89504e470d0a1a0a0000000d494844520000010c0000010c0802000000beee1fc400000009704859730000133900001339018fc256010000001174455874536f66747761726500536e6970617374655d17cedd0000200049444154789ced9d797013679afffb92ba75cb966df936c60736c6d8dc983b2181217742b25393c924339b99aaa9d4662b95daddaa546ded642ba9adcd647f331b26c9ce4e4298842c04420e0810480849b81cb0c118b08d6d2c1f92ef4377abd5e7ef8f37238cb18d2db5d492fc7efe32427af548ea6fbfeff3becf814a9284402090a9c1943600028977a04820903b00450281dc01426903e61c8220200882e3b820081cc76118a656ab1986110401c771954ac5711ccff3388e9324298a22c7710882a8542a8220445114451145511cc795fe1c7308143aeed1005ccd4009c16010c3300441689ac6308c20886030c8b2ac46a3c1308c61188ee30c06832449c16010fc8d2008c3300cc3e8743a201b9aa61104d16ab52a958ae7799665310c234992a228822080f0300c53a954ca7ef0a4048a4436589645519465598ee3803c02818046a3c1711c3c425114b89a49929424094c1d388e4b92248a228661188681c7c1df0882f03c2f49128661388e83911104c1300cbc512010080402188669b55a1cc7c1b403a44292a446a3415154e12f252980220907f0a50982c0f33c8aa23e9f0fc77170e15214a552a9044150abd5e01e1fbae2e535202400b06c1345114110144505416018261008f8fd7e8aa2341a8d56abd5e974b2db30778022991de072e4793e180c02274192249d4e278a22411004418cbf7c63cff85f5392249aa6fd7ebfd7ebf5f97c1a8d46ff37a04b332ba048ee0c58e4d0340d94c0711c49922a950a45518220126249234992ef6f8019c66030e8f57a9d4e471070f3e60e40914c0ef0bcc18c812008cff35aad16f806497055b9dd6eb7dbedf7fb398ed36ab5402d1a8d26093e5a348022b9054992789e076b2a8661d46ab54aa502fb48caaea3e445144596657d3edfc8c888c7e3e1388e2008a3d1683299b45a2d45516ab51afa3021a0487e84e3384992fc7e3f8220c0bb50abd5288a26fdb5e276bb070707bbbabadadada44514c4d4dcdcfcf2f2a2a4a494951abd5604342691b15664e8b44fa1b3e9f0f41104110f47a3d8aa273f6b4a1a7a7a7b5b5b5adadcded7603a9e4e7e7a7a7a7abd56aa54d5392b92b128ee3c0d9b62008244982136ea58d521e4110ec76fbb973e71a1b1bdd6e774141c1ba75eb0a0a0a0c060338ca4cfaa9f576e69c48c07982cfe7031f9c2449b0ac828b8a09dcb871e3db6fbf3d74e8505757577575f5ead5ababaaaa162c5860b158e69a4ee690482449621886e7799aa68d4623089482da980630ab1c3d7af4e8d1a30e87a3b8b878e9d2a54b972eada8a848494901c1044adb180be6844840bc462010008eb85eaf471004ca63e65cbe7cf9c081039f7ffef9f5ebd78b8b8bb76ddbb661c386eaeaeaacac2c8aa294b62eeae02fbffcb2d2364411b0d7190804789e07e18020a2092a645658add6cacacaecec6cbbdddedcdcdcdededed2d262b3d95c2e178aa25aad962088245e8325ed4c02826a19864151143a1e7241d3f45b6fbdb567cf9ee6e6669665f3f3f3376edcb86ddbb655ab5681592529bfe1249c4944510c0683344d0783418d4603a2cda1426441a5521516165aadd6fefe7e87c3e176bb3b3b3b9b9b9b6fdcb8e1f178542a9546a3493e4f2fd944120c06038100cbb20441180c063081286d5452613018162d5a545555e5f1789a9a9a58961d1c1c6c68686868681819192108c26432e974ba64fada934724822078bd5e90b9a1d3e9288a4ae255b2b2601896939373d75d77592c1687c3313232822088c7e3e9ecec6c6d6d75381cc160302525852088e438b04f069180bd5dafd78b611808d483f288011a8d66debc792449da6cb6b1b1310441188671381cb5b5b52d2d2dc1609024498bc59204a7f5892d1249925896a5695a10048aa2743add1cd9b98f075014351a8d9595955959598d8d8d4ea733f45f4ea7d366b3757575310c63341a298a4ae829258145028e0581fba1d16848924cdc9f2171214972f1e2c5cb962debeeeeeeeded05f9c98220389dcee6e6e6f3e7cf8f8e8e220862341a13373b325145c2308cdfef174551afd783c82ba52d9ad358add6bcbc3cbbdd1ed20980a6e99e9e9e8e8e0e9ee7ad562b98ea13ee5e966022013513fc7e3f48f630994c89f8a5271f044164656559add6f6f6f6dededef1ff0574d2d2d2d2d5d5c5715c6a6a6ac2c5b32492488007026213b55aad56ab55da22c84d542a556969696969695353535f5fdf84fff5783cd7ae5db3d96c9224994ca68c8c8c04bab52592480281004dd32449ea74ba24d833494a0a0a0ab2b3b3dbdbdb87878741c0750849929c4e676767e7c0c08056ab4d4b4b4b94c0fbc4108924491e8f876559b3d90c334be39cecececb4b4b4a6a626707e321e41104647479b9a9a5a5a5a388e331a8d068321fe53dce25d2420c6c4ebf5e2386e301812a53ac95c46ad56171414b02c5b5b5b0b0acd4c401084c1c1c1b6b6368ee3525252323333e3fcae17d722110481a6698661288a02a1a64a5b04991124491616167abddebababa499f208aa2cbe5eae9e9e9ebeb53abd5191919f11c1c19d722f17abd2ccb82d3a838bfd9402660341acbcbcb7b7b7bbbbaba789e9ff4391e8fa7adadadadad0dc77193c904f62a636ce74c884791807d5e8fc7238a2208018adb7b0c641a525252b2b3b3af5dbb36303030c1890f2108c2c8c848575717cbb2b9b9b9168b250e7febb8130908c402d5750d06437cde5a203324232343ad56fff0c30fa01ecda4006fdee1700c0d0da5a5a5592c9678bb2dc69d4842fbbca052bad2e64022822088dcdcdcfefefefafafa699e0676871b1b1b474747552a556666665cb928f125129fcf47d3b4d96c864e48d2a0d3e93232329a9b9bed76fbf4cf1445d16eb777757581c3fbf8392c8e17910882e0f7fb83c1a0c96482995249464e4e0ecff38d8d8d1e8f67fa67721c373636e6703844512c2929d16834b1b1707ae24224201c4b10049d4e072bc425250b162c686a6a6a6b6b9b6aa72b04c7710e87a3b3b393e7f9d4d454b3d9acf8aa5b7991f03cef743a41be145448b2a252a9cc66f399336740e4fcf44892343232d2dddd8d6198d56a4d49495156270a8b44922497cb254992d96c866785490cc8d072b95cf5f5f5779c4c001e8fa7afaf4f14c5828202b3d9aca08faaa4485896f5fbfd9224cdc1ca997310ad564b92646363e38458faa91045d1e9743a1c0e9fcf376fdebc949414a52e12c544c2f3bcd7eb45511476279b3b141414d86cb6ab57af82d64877046c0d034f26353535232343914b45196982f3238220f47a7dfc07814264e485175e484d4d9dd5eee5d8d8d8810307f6eddbd7dada3a3eed316628309388a2e876bb110489878d0b488c013b99df7ffffdac5ee5f3f9868686388e53c43f89b548388ea3691acc9e50217310922483c1e0952b5786868666fe2ab0eeeaefef0f0402c5c5c5269329963a89b548c071125c65cd65cc66736b6b6b4b4bcb0cb7b942b85caea6a62683c1909999999a9a1a339dc45424636363922481eaa3317b5348bc01e24dbef9e61bafd73bdbd7721c373434044a129b4ca62858370931d22268660bbae7c03904b26ad5aae2e2e2305e288a624747c7c71f7f7cf4e8d13034161eb198492449a2691a28642ef47c81dc119ee73d1ecf891327c27bedc0c0c0e8e8a856abcdc9c9d16ab5d18ef48b854840f4bb46a389c1e78124041445f9fdfeab57af0e0e0e863782c3e1f07abd292929050505d18e888dfa724b14c5b1b1319d4e07150219cf962d5bf2f2f222f14e2f5cb8b067cf9e4b972e4d5a6e4246a22b128ee3dc6e37c8a082812790f1f03cbf65cb169665c31ec1eff79f3f7ffea38f3e6a6d6d95d1b0db89624c2170454451349bcdd17b1748824210445555555151514747477823489234343474fcf8f1b4b4348aa2e6cf9f1fa593b728dedddd6e37481181db599049c9cccc5cb4685184837475757dfae9a7070f1e1c1c1c9caadc4484446b26e179dee572656565c12311d9090402dddddd288a0a82d0dada2a49526e6eeed2a54b132ed7202f2f2f3b3b5ba3d104028148c6b97efdfa471f7d949f9fffc0030f4423e9372a5f2bcbb25eafd76c36c3242a196159d6ed761f3b76ecbbefbe733a9d972e5d72bbdd56ab154551afd7ebf3f9962e5dba7efdfa4d9b36ad59b32621be798d46b361c3860f3ef820c27124496a6d6dddb3674f4141c1f2e5cb655f74c9dfa25a10048fc7038a87c3e82c590806836d6d6ddf7ffffddebd7b5b5a5abc5eef54011d1886994ca6d2d2d2eddbb7df7befbd0b172e8ce7995c14c54b972e6ddbb6edf6aac16160b55a7ffef39f3ff7dc73454545918f361ef9cf49c04f080bbfcb054dd3070e1cd8b163c7eeddbbdbdada18869966e50daa96f5f7f75fbf7edde170646464e4e6e6c6edbe228aa22e97abadad4d96ed29bfdfef76bb0d06434141815eaf97f3bc41921551143b3b3bc10f29efc87393a1a1a17ff9977f292d2d0de38e83e3786161e11ffff84730b1c7276eb7fba9a79e92ed6a4690f5ebd7efdbb70f94ff94cb483967128ee3bc5eaf4aa532180cf0dc304224491a191979f5d557df78e38d91919130928d244972b95ca74f9f76bbdd252525a9a9a9d1b033424892bc7cf972434343240726e31919194151b4b0b0303b3b5bae2954369180bb822449b036a92c747777efdab5eb0f7ff8c36ce3c927c0715c6f6faf56ab5db87061fc947b0bc1715c3018fceb5fff2ad7808220f87c3e8aa2162c5860341a651953b6d52acbb2814080a228782a220b9f7df6d9071f7c20cbfdd5e170ecdab5ebd4a9538ae4be4e8f4aa5d26ab5858585720d288a626f6fef37df7c73faf4e908779643c8231210a06530186080962cd4d5d5eddab5cb66b3c935604747c7cb2fbfdcdfdf2fd780324292644e4e8e8c030683c16bd7ae7dfef9e733a9853713641089288a344d4b92a4d3e9221f0de2743af7efdf7ff5ea55196ffca228da6cb6dffffef70cc3c835a65c5457574f53733e3cbc5eeff1e3c73ffbec3350d52dc2d1641009c330c160105606928bfafafa93274fca3eaccfe7abadad3d7ffebcec23474e4a4a8aec637abdde83070f9e3a756a86e58ba6410691b8dd6eb55a0da711b93879f264d8317fd3d3d4d4545f5f7fc7aad531c6e572451ec1753b2087f1f0e1c31d1d1d11c674452412499280bfaed3e9e0342217c78f1f8f52626a20103874e8905c9bad72919e9e1ea58eca3e9feff4e9d3a74e9d8af0fb8cc83296657d3e9f82f527938ffefefe8686862845b32208d2dbdb3bc32aa33143a552310c138d2e0b92247576761e3a74a8a3a323120f3ea28b1b0854afd7473208643cb5b5b5519d9359969dbee954ec9124293d3d3d4af39b2008e7ce9dfbf8e38f23994cc21789288a3e9f4fa3d1245c84763cd3d8d818d500dea1a1a1e84d53e1e1f7fb333232a29782ebf7fb8f1e3d7af2e4c9b077f6c214092855aad56a351a0d3c189191818181a8ae5d83c1a0cfe7937dcb3512341a8ddfef8fdead411084cececeafbefa2a6c0f3eccdf231008701c27d7b13f04100804b45a6d54af60bd5e3f383818572b641cc7dd6e77547747c1f6775d5d5d788bae7044224992cfe723082221327b12088661a21d86a852a9e27079ac56ab233fcd980649926edcb8f1f5d75ff7f7f78771b6188e48188601192370a1252f609f30aaf5fb58969531504a163c1e8fc160f0fbfd517d1796651b1a1a4e9c3811861a672d12106509ba16cdf6b5903b62b158a271fc1c826198e5cb97476ffc30200822063e9220087d7d7de7ce9dbb72e5ca6c3d93598b84a669500365b62f84cc84071e7820aaf7548bc552555515bdf1c380a2a8eeeeee181c46bbddeecb972f5fbc78716c6c6c562f9cb548fc7ebf5aad86f1f0514210844d9b36456970954a75f7dd774769f0b0c1306c68682836111b2d2d2da74f9feeeded9dd5d9e2ec44c2300c9846a0371225323333376edc1825f73d2d2dedde7bef8dc6c891303838886158cc8265befffefbbababa68890478231a8d067a23d18324c975ebd695979747e3b464f1e2c5ebd7af977dd808096fc7296c5c2ed737df7c73e3c68d99bfe92c7e894020208a22f446a24d7575f5c68d1b65ff9e8d46e3c30f3f5c525222efb091e3f3f99c4e67ccde2e10085cbb76edfcf9f333f7fd662112703602bd91684310c433cf3c53515121e39824496eddbaf5befbee93714cb9006188b17cbb9e9e9e1f7ef8a1b7b77786db5c33150988d482c5e163008661f3e7cf7ff5d557e5cab24051b4a6a6e697bffc654141812c03ca08cbb2344ddbedf658bea9c7e3a9afaf6f69699153242041972449d8a72a361004b179f3e6175f7cb1acac2cf2bb92d96cfee77ffee738f4461004e138aebdbd3df6efdbd5d5555b5b3bc3bde019fd001cc7b12c1bb33e8e10c0fdf7dfffdbdffe76e1c28561eb44ad56979797ffd77ffdd7962d5b0c0683bce6c9028aa28d8d8db17f5f97cb75e6cc99c6c6c699ecaacde8dba7691a4110388dc4988c8c8cdffce6372fbef86278a54a8d46e3962d5b7ef7bbdffdfddfff7d1cc66b010882502aedbeabababaeae6e7474f4cedb5c3329f3d8ddddedf3f9044190ab6e2464568c8e8efedbbffddbca952bcd66f34c4ea8b45a6d7e7efebffeebbfd6d5d5b12cabb4f9d31176cfc4c82108e2f1c71f6f6868e0797e7a23ef7c83e1388e61188aa2a0cbae1466b3f9d7bffe756565e5b163c73efffc7397cb258a228661a0dc2d8220288a621826080245518b172faeaaaa7afef9e77373734d26539cff6a67ce9c51eaad4551bc7af56a7373736565e5f4cfbc83482449f2fbfd5021ca8261585e5e5e5e5edee38f3ffee69b6f9e3d7bf6cc9933f5f5f5dddddd168b45a7d30583415114376cd8306fdebcd2d2d29a9a1aa022a50dbf03a2289e3e7d5ac1771f1a1aaaababbbebaebbb2b2b2a679e61d44220842201090b9903d2402288adabc79f3e6cd9b11041145b1bdbd9da6e9ecec6cabd50a529780fb11ff0a41108465590545822088c7e3696c6cbc76ed5a5a5ada34078077100958d1c655221b240486610b162c08fd33b1b61f455174381c172f5e54d00690d9dbdadaba62c58a69dadf4e77bf114531180caa542ad88e07223b7ebfffdb6fbf55da0ac46eb75fba7489611869ea3daee944c2b2ac2008d1288804818c8c8c28e8b5879024e9c2850bd3c7054f27128661701c87d308241a80630aa5ad404451f4fbfdd7af5f9f26c8723a91d0341d9f750320890ec3300d0d0d2d2d2d4a1b82200862b7dbaf5fbf3e3a3a3a5528d79422e1793e180c92249910fb2490c4a2bdbdfd871f7e50da8a1f1104a1b1b1d1e1704cd9d378aa570602812895318640eaeaeaae5dbba6b4153f82a2687f7fff340d61a6d4003c438444095114bffcf2cb1887c74f8324495eafb7bdbd7d787878d2274ca90190a90bcf1021f2228aa2dd6e3f70e0405c955af57abd6d6d6d1d1d1d9396249e5c24c02189c366ad904467707070cf9e3d4a5b31093d3d3d7d7d7d9346ce4f2e12bfdf0f8f4720d1a0b1b1f1c489134a5b3109bdbdbdcdcdcdb39849bc5e2f2cf8008906478f1ebd72e58ad2564c028661d7af5f9f34db7e7291f87c3e281288bc088260b3d9f6eddb373232a2b42d93e0f7fbfbfbfbfbfbfb6f5f714d22124110188681418d1079e9ebebdbbd7bf7d0d090d2864c494f4f4f5757d7eda5862611099c4620d1e0cc9933478e1c51da8ae9181818e8eaea02d988e31f9f5c24701a81c88b2008efbfff7e535393d2864c874aa56a6f6fbfbdf12a140924ea0402818f3efae8ebafbf061545e21696654747476f8f4f9928125114412a620c6d832433a228d6d7d7efdab52bde1a9ade0e48e81d1a1a9ad08274a248fc7e3f4cd685c8487f7fffeeddbb954dd39d390e87a3a7a7674210d74491d0341d9f55cc2009cafefdfb8f1e3d1ab3ce0a1182a2a8c3e1080402e31f9c2812d0003686564192169665ebeaeafee77ffee77657386ee138aeafafafa7a767fc83b7884492248661a04820b270f5ead53ffde94f8a94fa0d9b6030e876bbdd6ef7f8fea3b788044e2310b9181a1adab973e7975f7ea9b421b3c6e170389dcef16ec92d2209068330ae1112391cc7bdfbeebb9f7cf2497c46a04ccff0f0b0dd6e1f7f9e784bfe3ac7712449c2ad2d482478bdde6fbffdf6d5575f9de0fe260a14458d8e8ed2341d8a3bb96526114511c771281248d8f87cbe2fbffcf2b5d75e4b50852008323a3adadddd3d3c3c1c3ad8b929128ee3406114281248d89c387162c78e1de7ce9d53da90f001db57a0bc3278e4e6724b1445280f48249c3973e63ffee33fe2335d64e6a0286ab3d9028140a8c0c3449f04ee6e41c220180cb6b7b7bff4d24bf1506f2e4224490a06837d7d7da1cafc37975b388efbfd7ed85c17325bdc6ef7e1c3875f7cf1c50b172e286d8b3c6018e6f57a43115c3767128ee3244982f51a21b34214c583070ffee10f7fb876edda5475ab128ec1c1418fc7130aa5b9290970d62e4912f44c20334192249ee73ffcf0c3975f7ed9e170c47f90efcc71b95c2e972b24849b226159163a249019c2b26c4f4fcf9e3d7b76edda3521d2290930180c6eb73b5439e596dd2de8904066423018bc70e1c2db6fbf7decd83197cba5b439f20322b842c9ee3f8a441004954a0577812133e1d0a1436fbcf1c6c58b17272427250d81408061188fc72308028ee33f8a449224d0cd5559e320f10ccff32e97ebedb7df7efffdf76d369bd2e644178661542a1588e0fa51241886419f04320d434343172e5cd8b973e799336712316c71561004313434846118d8ecbd29129fcf374d6f45c85ca6b6b6f6e0c1839f7efa69777777a2e4184602cff31cc7391c8e9292929bcb2d9ee7310c839ddf20e391244910843ffde94f9f7efa697d7d7db27a2093a2d1684451e438eea648388ed3ebf58220c08624100441c07db4b1b1f1edb7df6e686848faf5d5edb8ddee50dac8cd2d608220a0e30ea1697a6868e8d8b163a74e9d3a75ea547f7fff1cbc2a50141d1e1ea6699a2449242412954ae5743a61b9adb98c288a8d8d8de7ce9d7be79d7706070747474727ed4330179024c96834fa7c3e8661288afa5124288a7abddeecec6c658d4b5c4211a30303037abd5ea7d3b5b7b7fbfd7eabd54a92a4c3e1e0382e3d3dbda0a0c0ed76fb7c3eb55a9d9e9eaeb4d50882208140c0e7f35db870e1f4e9d367cf9e8d87deeaf18046a30171374868260906837abd3ef44b43ee08684d3c3232027a8879bd5ee0da320cd3d6d6363434949e9ecef3bccd66e3386edebc792449badd6e97cb555454346fde3c93c9a452a9743a5d4e4ece82050b962c5912cb6f9e6559b55addd5d5d5dcdc6cb3d93efdf4d3f6f67697cb15572dda94c5e3f168b55a1cc79190484451a4280a3aeed303aeade6e66687c3313c3c7cead4a98181018fc70354613299bc5e2fd8219d7032dbd5d505fe4051f4faf5ebe06f8aa22c160b4551922489a2585959b962c58a356bd6ac5bb78e2449511441b0a95cbf08cff304418c8d8db5b7b78f8c8c7cf7dd7757ae5ce9eeeeeee9e949dc54dbe8012a028340ad1f45421004fca6a64292a4b1b1319bcdd6d9d979f8f0e1fefe7e9bcd66b7db310ce3793e141f3e3a3a1a7ac954ceeef81a1c0cc384aab681c281478e1c0175ffd7af5fbf76edda952b57565555e1382e49925aad369bcda1d3de3b066b8ba248d3b44aa5e238eec68d1b636363344d7ff7dd77edededa3a3a367cf9e2549727c6929c804727272689a0677969b8ebbcbe54a4b4b53d6b2b882655950a7ecf3cf3fbf72e54a5353d3a54b97288a8ac6714168f90b387dfa74a8786e5555557979797e7ebed96cc671dc6834e6e6e69695950583c19191111445737272b45a6d6767a7dfefb7582c82205cbd7ad5ebf58aa2d8dcdc6cb7db398eebe9e9b971e3864ea70b05ed41854c0f28f67b8b4fc2b22cd8ed82005a5a5aeaeaeaeaeaea8e1f3f3e3636e6743ac1e410fb03b5c6c6c6c6c646f0777a7aba56aba5697a7878382f2f2f2727074190bebebe60309897972708427b7bbb288a168b0514c5411084208890fc6eefe104998a6030a8d3e9c05af76614b04ea70b2dc2e62c8140c0e572eddebdfbabafbeba72e5cae8e8e8ed7d8f1464787838f4b7dd6eb7dbeda17f0e0e0e86fe1edf066442ab0dc80c71381c6ab51ac4a0dcf4491886899fab21c64892e4f7fbaf5cb972f6ecd97dfbf6b5b6b6c27d9e394e5a5adac4a42b8aa2ec767b4a4a8a72562986c7e3e9ececdcb367cff7df7fdfdcdcecf57a95b608a23c56ab559224508ceee639c9dcf4495a5a5a8e1f3ffed65b6f0d0c0cf8fdfe393b974226409264a82ef6cda42bbd5ecff33c383d497ac00ee9b163c776efde7de8d021a5cd81c41d3e9fcf62b1dc724ea252a95896a5284a51c36284d7eb6d6d6dddbb77efc183073b3a3a9436071277e0380ea24f6e49bac2719c6198d4d454456d8b05434343478e1c79f7dd772f5dba34a712242033471084b4b4349ee77f940a7894e7799065a2ac71d1666c6cecfffdbfffb773e7ceb1b131e87e40a6c16ab582d228374582a228885348d61517c3309d9d9dafbcf2ca9123473c1e8fd2e640e21dbd5e1f6a437d73b995c487eec3c3c35f7df5d5bbefbe5b5757074f9d2177c460301004a1d3e96e89024610044551104b9764a5b75c2ed77befbdb773e7ce8e8e8ea45f4f426441abd5eaf5fa495a2fe8743aafd79b4cc989a086fe8e1d3b5e7bedb5f1911a10c8f40882909191110ad1baa5cc2988534a8e994410841b376ebcfffefbefbdf71e54086456389dceb4b4b490f7714ba305b55a2d084272745fb871e3c61ffff8c7fdfbf73b9d4ea56d8124182693c96c3687447233eb8da228b031ac906172d2d3d3f3faebaf1f38702029cb3943a24d5151914aa50aa5d3dd9c34701c0f0402168b4521c364637474f48d37ded8b973a7d286401295ececec507f6a64bc480441d0ebf5c16030a12b02db6cb60f3ffcf0c30f3f54da104802939e9e6e369b43bb5b37975b188681b2d90a1926030cc3ecd9b3e7fdf7df1f1a1a52da1648a2a2d1688c46635a5a5a6877eb1691806e0c891bafb167cf9ebffce52fdddddd4a1b0249602c16cbf86904192f120441341a4d30184c4491300cf3c30f3ffce77ffea7dd6e4f9af696104530180c20702bf4c82d2251a9540cc324e239c9a54b97fefbbfffbbbdbd5d694320090f4992797979e343b46e118924491a8d26e1dc92fefefe0f3ef8e0c489134a1b024978701c57abd5858585a1b4446482484892140421e1ea6b7cf6d967fbf7ef1f5f1b0e02090f92240b0a0a288a1abf9e9afc703d51825344516c696979ebadb7e0b13a4416701c2f2a2a1a3f8d2013661204418c46a3d7eb4d0885200872fdfaf5b7df7ebbadad4d69432049822449656565138e0a278a04a46225ca06d1175f7c71ecd8b1845b1f42e216144517669ef20400001577494441542d5a34feb81db97db9a552a9701c07ade262685b380c0d0dedd8b16360604069432049028aa20b172e1c9f4902982812503a36fec31cfbfbfbdf79e79dbebe3ea50d812415555555b7a7544dd2fb42a3d1c47f1991dadada2fbef842692b20498524493312098aa2a00657ac0c0b078fc7f3f1c71fb7b4b4286d0824a9405174c99225339a49701c8ff30e2f172f5e3c7cf8302ce90091110cc3962d5b96929272bb373e8948d46af5f8064ef1465757d7c71f7f0cabbe43e4852088a54b9782de3d1398bc1f1f6823169f918eb5b5b5b0432c4476288a9a74ad854c2512d04e290e452249d2e1c3876d369bd28640920d83c1505d5d3de1840430b9484041c738dce3b2d96cd01b81c88e56abadaaaa4a4b4b9bb40aca94cb2de46f4586a26bdd6c1819193974e8102c520a911d93c934d55a0b994a24088268b5da786b1077e3c68d53a74e296d052409494f4fafacac9c74ad854c2392383c7a6f6868b876ed9ad256409290dcdcdcf2f2f209c1bf21a6140986615aad36100844cdb05973e6cc999e9e1ea5ad80241be9e9e98b162d9aca2141a6170908078e9315d7d5ab579b9b9be33c14009288e4e7e72f5cb8709a4a5a538a0441109224398e0b35ea55961f7ef8a1bfbf5f692b2049484949494949c9349d79a613094110288ac649b6c6a953a7dc6eb7d256409290458b16cd9b374fad564ff584e94482a2a8c160f0f97cf1b0e23a79f2643c8793411294c2c2c2f2f2f20985b626309d481004010d481577dfebebeb070707e341ab3384200893c9842088d56a5db264c982050ba6f90d200a52595939fd34824c550822844aa54251149c2a2a98f87ee6cc99b80db81c0f8661a228d6d4d4949696ce9b376fc3860da2280683419d4e67b158babbbbebebebcf9c3973f6ecd9388c669883a4a6a66edab4293737f70e79b8d29de0797e787898e7f93b3e337a3cf6d863b1fadec204c3b0949494cacacabffce52fe7cf9f0f0402e0ab0bb546124511fc93e7f991919177df7df7f1c71fcfcbcb83338c822c5bb6acb6b69665d9e92fbf3b8b441004a7d3093c13a5b05aad4a7f9fd3a152a9962f5ffeca2baf747777fbfdfe3b7e1c51149d4ee7f9f3e7fffddfff3d35355569f3e728a9a9a9cf3df75c7777f71d7faf3b8b4414459fcf373a3a1aba29c6982b57aec4795be09ffef4a75f7cf185dbed9eed471304a1bbbbfb57bffa5542b7bb4850962d5bf6de7bef8d8d8dddf167ba73e73790d01b0c0695eaf27ee1c285b82ddd4210c4d34f3ffd4ffff44f454545d33b7f938261586e6eee8b2fbe68341af7eedd0b3b46c4928a8a8a8a8a8a19dd9e6678cf733a9d5eaf5791c9e4f9e79f9f2af24c590c06c37df7dd37935bd1f40882d0d2d2f2ab5ffd4ae90f3487c071fcadb7de1a1e1e063125d33353910882e070386632a2ec6cdbb62d8c9b740cb8fffefbbffdf65bb96e1c57af5eddbc79737274758d7348925cb56ad5d9b36767783dcf626bc56030c43e97c36eb7d3341d87215b050505cf3cf34c4d4d8d5c3be38b162d7af5d557e7cf9f0ff7bba28dd56addba756b6161e10cbfea99fe1ee0545192a418877281cd8758bee30cb9efbefbb66edd2aef8e424545c533cf3c03279368535656b66ad5aa496b3e4cca2c6e5a20da3ec60522802f14b3b79b212525250f3ffcb0d1689477589d4eb771e3c6eaea6a7987854c60c58a1515151553658fdcce2c4482a22848578c7126561c06ffae5fbf7ef1e2c5b20f8b61587979f9ba75ebe27ccb3b71c171bcacac6cf9f2e556ab75e65ba6b35bfe8272da31f3105896c5302c0e4b62df7befbd696969d11839353575e3c68d999999d1181c62b158eeb9e79e850b17ceea30637622c1300c545189cd6482e3f8e0e0601c1e92dc7df7dde31b4fcacba64d9b1222502d11292b2bdbbc79735656d6ac5e35eb8d14b01b1b9b1a11388e777777cf7ced181baaaaaa2c164bf4c6f7fbfd0f3cf040f4c69fcbac5ab5aaaaaa6ab6f10db3160986611a8dc6eff7c76632a1693ade6692e5cb9747358637252525aa229c9be0385e5a5aba7af5eaccccccd95e51e16cc993244910044dd361bc7656d0346d3018e22d21b1b0b030aa5e19455120a6387a6f3107c9caca7aecb1c7aaababc308ad0a4724288a6a341a9ee7a39dd9cbb26c1c9eb51304113d8704909a9a0a134ee4a5b2b272d3a64d56ab358cc3df300f77d56a358ee3d1fe21b55aadcbe58aabc335bd5e3f38383855a93f59f0f97cf156cc29d151abd560d73e3cff364c91806dae68d75251abd5344dc795e3eef3f97c3e5f54b3ede373d73b71d1e974ebd7af5fbb76add56a0d2fe427fc30219224310c8b6af1798ee3ac566bbcb522114531aa31ed2449363636466ffcb9c6fcf9f31f79e491850b17861d1417512c9d5eafe7793e7a8b2e9ee7d3d3d3a33478d8d4d7d747d571c771bcb6b6367ae3cf2950145db162454d4dcdcc23b56e272291e038aed56a59968d9207afd16894caf49a86dededededede280d2e8a624343431c46e22422288a9694943cf4d043e5e5e591ec00451a954d51148aa220953ec2a12645ad56c7dba101cbb2870e1d8a9263edf57a8f1f3f1e8d91e720595959dbb76f5fb1628556ab8d24a3215291a0284a5154f42693458b16399dce688c1c361e8fa7bebebea9a9291a837774749c3c79321a23cf35300c5bbc78f1a64d9bcc6673a443456e0d887a8c52f72993c914ed438930387ffe7c6d6d6d3476144e9e3c595f5f2ffbb07390458b16fdf4a73f0deff470023288049c2d4a92148d8bc6e7f34523283d421886f9fdef7f6fb3d9648ccd1104e1ead5abfbf7ef8fb799331151a9549b376fbeebaebbd2d2d222cff49427539420088d46c3b2acec0708b9b9b97178e88e20487f7fffcb2fbfdcd1d121d780369b6dd7ae5d757575720d3867d1e97435353577df7d776666a62cb9d0b2a5535314a556ab19869137cc3b3737d76834c6a14e04413875ead4ae5dbb1c0e47e4a3b12cfbc1071f7cf6d967910f05292b2b7beaa9a756af5e2dd7652367cd01b0e892f70cc1e7f3ad5ab52a0e0b412008323a3aba7ffffe4f3ef924c2584f4110fefce73fffeffffe6f7777b75cb6cd59525353efbaebae75ebd6994c26d9aa57cb520e07208a22c330636363f2561edabb776f3cd70235180c2fbcf082cd660ba35c7230186c6d6d7dedb5d7e2f0cc341151a9540f3ef8e0b973e7e4bd02e514892449a228d2343d3a3a2a6381eda3478fc6a1ef3e1e8220366ddaf47ffff77fdddddd33ff79464747bff8e28b5ffce217b0c6a92c6018565151b173e74e97cb25d7b5079039c016ec74711c170804b45aad2c6ed3aa55ab743a1d8ee3719bd4caf3fcb973e75c2ed7e5cb97b76fdf5e5a5a6a3299a6faec3ccf7bbdde969696bd7bf77efffdf76d6d6db039912c9495953dfdf4d31b376e943d461b95a270522e4992dbed2649129cc747381a4dd3fff00fffb06bd72e596c8b362449ae59b366d3a64debd6ad5bb66c99cfe7f37abd288a5aad56b7db7de9d225a7d3f9dd77df1d3b76cced76c7a7af958868b5daa79f7efab7bffd6d841128931215912008120c060381804ea78bfc289061980f3ffcf085175e88d27965f4200862e9d2a5b9b9b99224b5b6b60a8280a268777737cc1591178aa2eebaebae975e7aa9a6a6262ad947f2aede420882e0f3f9c6c6c66429957be4c8918a8a0af93f3c24f1a128aaa6a6e6af7ffd6be495cba7225a6567310ca3280ac771508b3ec2d12a2a2a32323264310c92642c58b0e0b1c71edbb06143e4315a5311c5dacc388e1b8d464992028140843a29282858b972651c06714194252d2ded273ff9c9430f3d949b9b1bbd9e9e512f601e4acc9222707e789ebfe79e7b60b975480814454d26d3c30f3ffcf8e38f17171747f5061af5cb0ec771922483c1602493094110595959f7dc738f8c8641129a9494940d1b363cfae8a391e4e5ce9058dc9bc15eb0dbed8e24e7a4a2a262c99225b01a15044110ad565b5353f39bdffc66f5ead531a813120b918013468aa27c3e5f24d5551e7ae8219d4ea7603b79489c505555f5c4134fac5dbb3625252506d743ec56f95aad962088402010f67c929f9fbf66cd9a487c1b48a24351d492254b9e7aeaa9071f7c30353535366e6a4c5d618d4683e37830180cef4237994cdbb66d837bc173161cc7172c58f0e4934f6eddba359631af3115098ee32037cbe57285f1728aa256ae5cb972e54ab8e29a9b545454fcfce73f7fe8a187f2f3f363f9beb1de54c5300c5440f2783c61042c2e5cb870fdfaf52693290aa641e29a8c8c8c471f7d74fbf6edf3e7cf8ff58959944ef2a74714458fc7e3f57a671b512f8a625f5f5f6161219c4ce60e2a95aab4b4f477bffb9ddd6e97310563e628733c070a1189a2385bff04455190e4146f4d4b20d1a3a4a4e489279e78f8e187c3682d220b8a9d61ab542a503bdde3f1cc4a277abd7ec3860d1b366c889e6d90380145d18a8a8a9ffdec674f3ef96445458552fd05940cf420080204a579bdde59f927e5e5e5f7df7f3fdce64a6e701cb75aad8f3ffef8934f3e595656a6643190d8aff0260082ea67eb9f7474746cddba152eba9215ad56bb74e9d2d75f7fbdb3b393e3b8e85d7e3341f99041d0ea449224bfdf3ff3f9242727e7d9679fd5e97451b50da2086ab5bababafa17bff8c5a38f3e9a9797a7781727e545822008411006834192a4b1b1b119d6de264972eddab5bffce52fa36f1d24d6ac5dbbf6d7bffef5f6eddbe7cd9b17178b056527b209300ce372b9689a9e613e635d5ddd9a356b94fe0a21f2806198c964bafffefbf7eddb07b673e284b8984942a8d56a954a150c066718dfb57cf9f2679f7d362b2b0b1e9b240156ab75ebd6adcf3efbec3df7dc1349cf1df9515aa51311453110080c0e0ebaddee99d4b01a1c1c7ceeb9e7a07392e8141717ffe33ffee3e9d3a7dd6eb72c75116424ee44021004c1ed7683a23bd33f93e3b8fafafa2d5bb628fd2b43c204c3b0eaeaea575e79a5b5b59565d978538814b722912489e3389fcfe7f178eeb8351c0c060f1f3e5c5656a6f4cf0d991d200577eddab56fbcf1464747476caeab30885f914892248aa2cbe5eaefeff7fbfd777cf23befbc13e3e0504884e4e4e4fcdddffddd279f7cd2d7d7a74850d60c8956713a19e179dee7f3e1384e51d434e19f344dbff6da6b6fbef9a6d3e98cff0f05a9acacbcfffefb1f7becb18a8a0a8aa2e2b9ca87c2c734338120089d4ec7304c2010c0306caa8d738d46f3e4934fdaedf603070e78bdde181b099921188669349a952b573ef1c413f7dd775f5e5e5e3ccbe347949eca660acff31e8fa7afaf6f1a2f85e3b8cb972f3ffae8a34a7fa990c9a128aab8b8f899679ef9f4d34f070606627c09854d022cb7c6234992cfe793248924499224277d4e4343c34b2fbd74e2c489b8ad423f37311a8d8b172fbee79e7bb66ddbb670e1c244aae9a1b048670fcbb26eb77b64640414bc9b94868686eddbb7cb5e821f121e388ea7a4a4fcec673ffbe8a38f1c0e07cff371b8cf3b0d09e0934c40a55211040112e50982d06ab5b73778a8a8a878fef9e7fd7eff77df7dc7308c52a6425014d5ebf5c5c5c58f3cf2c8238f3c52525212833259b29360cbadf180ee733ccfabd5eadb7b75b32c7bfefcf9975f7ef9e4c9938a9807c171bcb0b070fdfaf50f3ef8e0faf5ebcd66b3e2f1bce191c0224110042400fb7c3e9ee78d46a346a399b055d2d5d5f5ca2baf7cf6d967b0397a2cc1715cabd5ae59b3e6e1871fdeb06143515111499209e381dc46628b2404cbb281400045d1db1d7a9bcdf6ce3befecddbbd76eb747de0402323d288a6ab5daf9f3e7df77df7d4f3cf1447171b1c16048804dde694912914892c4b22c4dd3c16050a3d1e874baf1337b7b7bfbeeddbbfffce73f0f0f0f2b68647283e3b84aa59a3f7ffe9a356b7ef2939fdc7df7dd06832141d757134812918400c58a789e0753caf8c4e883070fbefefaebe7cf9f8fa46e3764528c46636666e692254bb66cd9b26ad5aaa2a2a2dbbdc4c425d94482fccda16718461445b0f705a67b4992ce9f3fffeebbef1e3f7ebcb7b737f93eb852949494545555ad5dbb76ebd6add9d9d949b0be9a40128a04204912c3303e9f0f41109224351a8d4aa56259b6adadeda38f3e7af3cd37dd6eb7d23626362449a6a4a4ac5bb76efdfaf52b56ac58b26449427be7d390b422010882109a5540fb071cc751146d6868d8b163c7575f7d35303000bdf9d9a2d3e9727272162d5a74f7dd77af5ebd3a3f3fdf62b124d9ec319e241749088661fc7ebf288ae050852088b6b6b6c3870fefdbb7efe2c58b4a5b9730188dc69c9c9c65cb96ad5bb7aea6a6a6b4b4942088e4f0cea761ae8804f9dbac42d334c3306ab5da6030a8d5eabababafdfbf71f3972c466b3c158af494151144551b3d95c5454b47af5ea9a9a9aeaeaeabcbcbc440abe8a8c39249210208d9ee3381445c1f6576d6dedfefdfbbffdf6dbb6b636a5ad8b2f0c06437a7a7a7979f9d2a54b57ad5a55555565b55a711c4fe2c5d5edcc45912008228a22c771344d830e752a95caebf53636361e3b76ececd9b33d3d3d91b4ad4b02300cb35aadf9f9f9d5d5d5cb972fafaaaacacece4e4f4f57b2d6a872cc51918c87e779700a29495230186c6c6cfcfaebaf4f9e3c79edda35a54d8b35388e6764649024b962c58aa54b9756565656575767666622088261d81c595cdd0e140982208824498220b02c0b5a69fbfd7ebbddded8d878f6ecd9a6a626bbdd3e3636a6b48d5144a7d35114959a9aba74e9d2050b162c5ab468cd9a351445198dc658b7cb894ba0486e4114459665817fef76bb1d0ec7f5ebd71b1a1a2e5ebc383030e0f17840cd1ba5cd9401b55a8da2686a6a6a4141417171715151d1f2e5cbf3f3f3d3d2d2d2d3d3098298b3f3c6ed40914c09cff38140c0eff7fb7cbed1d1d12b57ae343434d4d5d55dbc783141bf341cc72549d2683425252566b3b9bcbcbca2a2a2a8a868debc797979792051674e79e433048a6446f03c0fd4e2f57a9d4e674343c3e5cb971b1a1a1a1a1a445104375d90c5a6b4a5b700820e5996cdc8c8c8cccc341a8da5a5a54545452525255959590b172e046519e09a6a7aa048660d68a812e2f2e5cb8d8d8d57af5e6d6a6a0282015e4d301854c43ca3d1c8f3bc46a3c9cfcfd76ab5595959797979c5c5c5191919858585f3e6cd4310049c1129625e2202451211400f344d07020186615a5a5adadada7a7b7b5d2e9720084ea7b3bfbfbfbbbbdbe572511445922488e70ffbed300c031e914aa552abd57ebf5fa3d1808e5fc160b0acaccc603058add6f4f4f4bcbcbc9c9c1cbd5e5f56564692a44aa5d26834e0681c3a1bb3058a443624490a0402c16090e33851148786864647473d1e4f201070b95c232323c3c3c32e978be77986618687877b7b7b419d4f5114c7c6c67a7a7ab2b3b3f57a3dc330fdfdfd66b339353595e7f9c1c1418bc5929696c6719c2008a9a9a9269349ad56eb743ab3d96cb1584c26935eafcfcfcf379bcd388ea7a7a75b2c16108033554119c8ac80228916402a0882e0384ed334cbb2822088a208bc1a86610882c0717c787818b4c253abd59224815464707d035550149591911108044451b45aad5aad9665598d4663b158489214048124499d4e0726190cc3a0e72d3bff1f247c807b352c43190000000049454e44ae426082")
var ResourceLOGOPng = &fyne.StaticResource{
	StaticName:    "LOGO.png",
	StaticContent: Logo,
}