package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testPlayerStr = `
{
  "data": [
    {
      "type": "player",
      "id": "account.c0e530e9b7244b358def282782f893af",
      "attributes": {
        "createdAt": "2020-07-01T14:38:28Z",
        "updatedAt": "2020-07-01T14:38:28Z",
        "patchVersion": "",
        "name": "WackyJacky101",
        "stats": null,
        "titleId": "bluehole-pubg",
        "shardId": "steam"
      },
      "relationships": {
        "assets": {
          "data": []
        },
        "matches": {
          "data": [
            {
              "type": "match",
              "id": "4b1ea7d9-b512-482b-a025-ae6ee5403c64"
            },
            {
              "type": "match",
              "id": "9bf1e80c-c3fb-49ec-8a5d-11a149125fbc"
            },
            {
              "type": "match",
              "id": "5b4b19f3-6858-484b-b04c-79317587f0f8"
            },
            {
              "type": "match",
              "id": "8c351802-97d1-4086-80a9-a0c706ebcccf"
            },
            {
              "type": "match",
              "id": "41a3b01f-f5e7-47c6-832a-5552a7a56975"
            },
            {
              "type": "match",
              "id": "51ad2b90-0f5a-4977-9024-de10624de047"
            },
            {
              "type": "match",
              "id": "793f33d3-1f65-423e-9c57-84d7a115ae8c"
            },
            {
              "type": "match",
              "id": "1ed21904-cb66-4c17-ad04-b9eb3e78801d"
            },
            {
              "type": "match",
              "id": "cf74ddd8-c39b-4da8-a7a3-a98f85a772a6"
            },
            {
              "type": "match",
              "id": "e6242e80-c6fd-49a6-a4b2-da0b899d5a4d"
            },
            {
              "type": "match",
              "id": "957b4fdc-11e8-4e54-b465-3d4a343589b9"
            },
            {
              "type": "match",
              "id": "e7cdfeb0-bb8e-4874-a62d-c1ce19f4269a"
            },
            {
              "type": "match",
              "id": "d772aa4f-6069-400b-8b9f-cabdfaf79069"
            },
            {
              "type": "match",
              "id": "bdb93d93-171f-48f8-8841-d3e45fc8a5ad"
            },
            {
              "type": "match",
              "id": "849480ab-3f3e-4e32-99bb-f11357f562a6"
            },
            {
              "type": "match",
              "id": "8fee1153-6dc9-4766-99d4-c6d9c2789451"
            },
            {
              "type": "match",
              "id": "1e1b4c93-2a80-4ab2-bdd5-ce5afe379da2"
            },
            {
              "type": "match",
              "id": "33a16aaa-840e-4db2-a9e5-5e26b89c88e6"
            },
            {
              "type": "match",
              "id": "0101a02e-26b6-4648-bfeb-f0aba69b1350"
            },
            {
              "type": "match",
              "id": "6a0ff82b-6d8d-47d1-a110-24ee6fe5b5bb"
            },
            {
              "type": "match",
              "id": "572e8004-3b16-40d6-af0a-c34cf5fd5298"
            },
            {
              "type": "match",
              "id": "97f59a35-2b78-4a88-b807-871791a901c3"
            },
            {
              "type": "match",
              "id": "5b3394b5-f71b-4c65-ba0a-858532510d69"
            },
            {
              "type": "match",
              "id": "070f4224-8bbd-43af-82b4-ae43696696d1"
            },
            {
              "type": "match",
              "id": "435ed36b-2b49-4d6e-8f28-e0e230257803"
            },
            {
              "type": "match",
              "id": "f4936ee0-ad78-4404-bf62-f7441fb40932"
            },
            {
              "type": "match",
              "id": "19738983-bfb1-4bd6-8804-305e0a34154a"
            },
            {
              "type": "match",
              "id": "64db0c84-65e3-4aed-83db-9257f6a68e6e"
            },
            {
              "type": "match",
              "id": "17e39928-1778-4810-90fe-de16294f41bf"
            },
            {
              "type": "match",
              "id": "6d7384bb-b42d-4c17-8ee8-f09f90ba3cd1"
            },
            {
              "type": "match",
              "id": "b0100198-c363-4366-b2ee-acec723598ba"
            },
            {
              "type": "match",
              "id": "1742f2be-5b2c-4cb1-906a-083bc4e4ea26"
            },
            {
              "type": "match",
              "id": "5b06bc1b-0d58-41f7-9639-e8aca8a1a61f"
            },
            {
              "type": "match",
              "id": "af8397cf-e86e-4975-9f65-202bc85985c3"
            },
            {
              "type": "match",
              "id": "ef665b14-50b2-472a-b9f5-670757be6f2f"
            },
            {
              "type": "match",
              "id": "55fe3200-4657-45e3-a7b9-1df909072ed1"
            },
            {
              "type": "match",
              "id": "acd114c5-1325-459e-ad97-5b9dabe7dc2a"
            },
            {
              "type": "match",
              "id": "6e3bb3e6-1815-475a-8446-23a83f4e6dcc"
            },
            {
              "type": "match",
              "id": "4517e048-20a2-4526-8217-36544550bb4f"
            },
            {
              "type": "match",
              "id": "7111b3f8-c240-459a-934c-b46e5213bda2"
            },
            {
              "type": "match",
              "id": "80cd36a4-22f2-4b4b-8ce2-9d133d89bb30"
            },
            {
              "type": "match",
              "id": "1daca130-6f2c-42bb-91be-9d6111bd854b"
            },
            {
              "type": "match",
              "id": "8e0f7ace-7d9b-4467-9681-f78e7bde7a6f"
            },
            {
              "type": "match",
              "id": "5ff6ea5a-0adc-46b7-8da9-af3cc1468c55"
            },
            {
              "type": "match",
              "id": "b3a46919-b894-4373-8499-f782998ca388"
            },
            {
              "type": "match",
              "id": "2c0ed7a0-5ca9-4a3f-8700-02d76e829e67"
            },
            {
              "type": "match",
              "id": "6334aef2-682f-4a5d-acad-aba1f5b0c5b5"
            },
            {
              "type": "match",
              "id": "8d5cf197-a062-45c1-9b0b-f71cb03c977a"
            },
            {
              "type": "match",
              "id": "02b9d148-c6db-451d-ae31-e4bf8e46d9fd"
            },
            {
              "type": "match",
              "id": "3d18df85-6ef0-4e50-8012-cb9468335afc"
            },
            {
              "type": "match",
              "id": "3e10da27-4310-4dbb-9fa1-e7c159526294"
            },
            {
              "type": "match",
              "id": "5abfeb04-b2c1-4ba2-82e2-cc5355c20402"
            },
            {
              "type": "match",
              "id": "e65db083-f41d-4c52-a0e0-a02e54d6cbd2"
            },
            {
              "type": "match",
              "id": "67ad30c8-eaa1-424a-b055-0e9dab5850c5"
            },
            {
              "type": "match",
              "id": "7ab68309-ce70-42e0-9e08-8a6f87a75880"
            },
            {
              "type": "match",
              "id": "1a705e72-cf27-41a0-a17a-2c4f35ab7211"
            },
            {
              "type": "match",
              "id": "9f37b15a-ac3b-471f-9d3a-935fd9472054"
            },
            {
              "type": "match",
              "id": "acd2ffaa-4b2d-4956-95ca-6c4e4f79a6c6"
            },
            {
              "type": "match",
              "id": "e69a90a4-35bc-4835-a7b2-cc07f5f5fe06"
            },
            {
              "type": "match",
              "id": "b8ef14d6-de68-4609-a066-7dec341f1508"
            },
            {
              "type": "match",
              "id": "afbb8a29-607d-43e5-bb73-418111fafc90"
            },
            {
              "type": "match",
              "id": "5bb0cdfc-745d-49e2-99e7-f23c82f4d794"
            },
            {
              "type": "match",
              "id": "13a93904-d484-47bf-a8b9-309c91f98f87"
            },
            {
              "type": "match",
              "id": "fb7f90cd-6816-4fcd-a357-aaae75e1171b"
            },
            {
              "type": "match",
              "id": "54d78194-2fd0-4bbc-b5b9-a2f1c7838b79"
            },
            {
              "type": "match",
              "id": "2d64c283-3778-4a2b-a71e-b1476d08a367"
            },
            {
              "type": "match",
              "id": "b6f01b17-40db-4801-ae74-12ce18bdb502"
            },
            {
              "type": "match",
              "id": "f8b5b5f1-5de2-4084-99cb-569a7b04de57"
            },
            {
              "type": "match",
              "id": "5335ad06-47b5-4c6d-bc45-ddeb309df924"
            },
            {
              "type": "match",
              "id": "34b37126-55f6-4ad0-b165-2cb3dc0791e6"
            },
            {
              "type": "match",
              "id": "91a20280-1bd9-4d28-97af-c000b5350bba"
            },
            {
              "type": "match",
              "id": "e8ae623c-2363-419f-b4ff-9e35c6f1683f"
            },
            {
              "type": "match",
              "id": "0c66bd20-2219-41fb-adb3-90d3effd23ae"
            },
            {
              "type": "match",
              "id": "14cc00e3-a04f-4cfc-a911-178675524cb1"
            },
            {
              "type": "match",
              "id": "b8daaec2-c8ec-480f-a942-963583663aee"
            },
            {
              "type": "match",
              "id": "b6e396fc-f887-41b3-8a43-3b894aa80c4f"
            },
            {
              "type": "match",
              "id": "90fd3947-ac48-4934-a2ef-1ad95642b7d7"
            },
            {
              "type": "match",
              "id": "dfc6b50c-f5ed-4f6e-a42a-a5f1a80ac532"
            },
            {
              "type": "match",
              "id": "5520c34c-e6df-4973-a658-6cfb1b18e804"
            },
            {
              "type": "match",
              "id": "2cf28a85-5538-4708-b1f7-87da6ff5c7e8"
            },
            {
              "type": "match",
              "id": "189cc9c8-4644-493e-8591-31b17eda9827"
            },
            {
              "type": "match",
              "id": "a8806cf7-982e-4772-9c73-7cec86f2c6ad"
            },
            {
              "type": "match",
              "id": "8eeea807-e668-4b97-911c-108363b8d850"
            },
            {
              "type": "match",
              "id": "f6d9f1d9-9421-41c7-a4f9-299068e2b41e"
            },
            {
              "type": "match",
              "id": "2ea35ad2-eaf5-4bc7-87e7-84ac4f968d5f"
            },
            {
              "type": "match",
              "id": "deb8ddef-219f-4dfa-886f-8039ddf4c1c7"
            },
            {
              "type": "match",
              "id": "1ed4f965-2876-4c35-96d0-f91086c7509a"
            },
            {
              "type": "match",
              "id": "6b364cce-db1d-47a4-83a4-a71101fada69"
            },
            {
              "type": "match",
              "id": "8be546c2-1c49-4812-a3a6-38d7fae4688e"
            },
            {
              "type": "match",
              "id": "3f387e81-4516-4e67-a3da-a8fca336b173"
            },
            {
              "type": "match",
              "id": "b09bb69d-f819-4b34-a4c8-e309ab91cbb1"
            },
            {
              "type": "match",
              "id": "2a178688-c7a2-456c-95f2-b4fed73b1080"
            },
            {
              "type": "match",
              "id": "77d4acb3-4ea9-4075-a70b-5fcd5bdb3bc3"
            },
            {
              "type": "match",
              "id": "fc65d1c9-ea0b-4767-9f37-14d5f5d6e283"
            },
            {
              "type": "match",
              "id": "bf4c4c18-926a-472d-92b1-af7ac67048c2"
            },
            {
              "type": "match",
              "id": "10aa121a-cf00-4e6b-9d38-1276404025ab"
            },
            {
              "type": "match",
              "id": "8391f602-3eba-4655-bfcc-7234f38def03"
            },
            {
              "type": "match",
              "id": "a4c7304f-e3e7-4a86-ad16-0076b4abb186"
            },
            {
              "type": "match",
              "id": "5b276901-ed8a-4919-bfbf-2259083204a0"
            },
            {
              "type": "match",
              "id": "27ae1a15-b88b-4c59-8a1b-3dfcd7ae0750"
            },
            {
              "type": "match",
              "id": "76761ab5-ed91-4e13-bf85-a80a1ef09ff0"
            },
            {
              "type": "match",
              "id": "9e9cde6b-5a27-4181-8360-13b41784b0e9"
            },
            {
              "type": "match",
              "id": "bc7dc2e3-8511-4e4d-909f-a6fc19ae2307"
            },
            {
              "type": "match",
              "id": "1626575a-7317-47a6-86e2-c9b21fca6e00"
            },
            {
              "type": "match",
              "id": "8662b9e1-a134-4dd3-9635-c8206bda36e8"
            },
            {
              "type": "match",
              "id": "8e969016-5857-4fe3-93eb-99bb6fd8f285"
            },
            {
              "type": "match",
              "id": "9aac3be8-5a0e-4270-a624-2360ca487ba7"
            },
            {
              "type": "match",
              "id": "3e16c1ad-1440-4322-af67-d7e475d31ad0"
            },
            {
              "type": "match",
              "id": "6c473717-df7d-49ef-8a82-21d2229af20c"
            },
            {
              "type": "match",
              "id": "86469d09-6d44-4956-8f8b-dd015e7d562b"
            },
            {
              "type": "match",
              "id": "baf00942-e898-4b55-9939-601464594847"
            },
            {
              "type": "match",
              "id": "55a6e92b-5f35-415e-a2d9-c9bb859bdaae"
            },
            {
              "type": "match",
              "id": "583e7f99-10d8-4ad6-8071-3ff8a84ad44d"
            },
            {
              "type": "match",
              "id": "589128c1-2b81-4dc7-b261-33766e228349"
            },
            {
              "type": "match",
              "id": "27dc3fbc-62a0-4b9e-8bdb-355a7960fc6c"
            },
            {
              "type": "match",
              "id": "338b6c8d-d360-4318-8e63-299eb372b563"
            },
            {
              "type": "match",
              "id": "f4b76e38-f15a-4acb-b92d-d21633ffe727"
            },
            {
              "type": "match",
              "id": "c7927e36-8f85-4db0-9426-1e1f52cda757"
            },
            {
              "type": "match",
              "id": "e267a861-4437-49de-901a-144cc036c8e0"
            },
            {
              "type": "match",
              "id": "e55c8123-3f11-49e8-af0b-2415e2174972"
            },
            {
              "type": "match",
              "id": "8fc428f2-2637-41af-8357-2d1211c675de"
            },
            {
              "type": "match",
              "id": "bc47f218-cfa8-4e29-be81-7c80678239c0"
            },
            {
              "type": "match",
              "id": "4ebc562f-20b0-4643-8d71-f80fa0f8d910"
            },
            {
              "type": "match",
              "id": "d47084d1-a451-4d98-a844-c0ed96e95020"
            },
            {
              "type": "match",
              "id": "e4b9e230-3db0-45f3-8a66-5960d9b17a61"
            },
            {
              "type": "match",
              "id": "2f588ec6-5ad3-43bf-a06a-eb5239f9f007"
            },
            {
              "type": "match",
              "id": "911d0ae0-2231-4d64-936f-1cc0181e5c21"
            },
            {
              "type": "match",
              "id": "4fa67911-a898-4cd5-ad30-0fc007241b1d"
            },
            {
              "type": "match",
              "id": "bb6ec23a-4eb2-418a-80b0-c9ae44fffd3f"
            },
            {
              "type": "match",
              "id": "117ad0e7-912c-43e0-a181-898e3d5068e0"
            },
            {
              "type": "match",
              "id": "c81f67cf-426a-4eac-bb77-37a130afe374"
            },
            {
              "type": "match",
              "id": "aee68df7-097b-43fa-b84e-3d4239fb5bc5"
            },
            {
              "type": "match",
              "id": "ad8b98a4-273b-4e65-982a-cc368d95b67e"
            },
            {
              "type": "match",
              "id": "c47961b5-50e6-41e0-8902-de6817c27597"
            },
            {
              "type": "match",
              "id": "63b9261e-e2e9-4c68-99fa-1f9b3c4c0ad9"
            },
            {
              "type": "match",
              "id": "360c6ce7-c2bf-4e73-a16e-9940e935a2ef"
            },
            {
              "type": "match",
              "id": "eeb0d614-2539-41ba-8da1-36414a36162b"
            },
            {
              "type": "match",
              "id": "a9baac9d-44be-48a2-b32e-96c641cb5ce7"
            },
            {
              "type": "match",
              "id": "72d2bb85-f49e-4843-b91b-8aa86584dcd0"
            },
            {
              "type": "match",
              "id": "bb0151aa-17c1-41fd-a368-c0c812f8781b"
            },
            {
              "type": "match",
              "id": "c539d85f-7107-4094-aed2-94142e96a2ea"
            },
            {
              "type": "match",
              "id": "40fda3ab-5813-496d-b20a-389d646d150a"
            },
            {
              "type": "match",
              "id": "9175161e-d710-4bd5-a4fe-78aaec36652e"
            },
            {
              "type": "match",
              "id": "2f50fc7a-a276-493a-87d6-a9a7cc48b00a"
            },
            {
              "type": "match",
              "id": "3f8c0046-f809-4906-a3ff-32e5adf79941"
            },
            {
              "type": "match",
              "id": "da09a7a5-8f2d-41b6-b415-7b51a66efdcc"
            },
            {
              "type": "match",
              "id": "5d426c01-e0d2-4d65-b536-3cf3d09d34ee"
            },
            {
              "type": "match",
              "id": "f03b1d61-d54b-499e-980d-ff925124865c"
            },
            {
              "type": "match",
              "id": "98dbb242-5ee6-40b0-804d-e82813058a1e"
            },
            {
              "type": "match",
              "id": "5ceda2d2-67d6-48ca-bd53-7d367c526c67"
            },
            {
              "type": "match",
              "id": "3c873f95-c682-485a-b25d-2210b42c5914"
            },
            {
              "type": "match",
              "id": "bf91422f-5455-49d5-bbe2-b1b3b9de78e9"
            },
            {
              "type": "match",
              "id": "b205e1ca-5e66-4c44-ad07-f9da8af255dc"
            },
            {
              "type": "match",
              "id": "3b71d319-8455-47b2-97dd-8de97eba4fc2"
            },
            {
              "type": "match",
              "id": "08ca75e3-d2d8-4ae5-96e0-91603ded505f"
            },
            {
              "type": "match",
              "id": "7d4b9634-d261-4b3a-9062-0e8c57ea65cd"
            },
            {
              "type": "match",
              "id": "13dcf6c8-1d18-4dff-b601-2c325bbea7f2"
            },
            {
              "type": "match",
              "id": "5a3d9931-a062-4cca-a408-b19fc92041dc"
            },
            {
              "type": "match",
              "id": "6a7a64ff-b5b4-4273-b889-b176be74f743"
            },
            {
              "type": "match",
              "id": "4806ae30-89e8-4eba-b9d1-5c78a8335968"
            },
            {
              "type": "match",
              "id": "d0961482-990e-46b9-b4df-518411b0cdc7"
            },
            {
              "type": "match",
              "id": "2afb8d77-e5cf-421b-93e2-98609c841701"
            },
            {
              "type": "match",
              "id": "967fc627-7059-4730-bf20-1edc22be51d4"
            },
            {
              "type": "match",
              "id": "19a6493b-0c5e-4217-aa9f-36c5ff7a60a1"
            },
            {
              "type": "match",
              "id": "cb4ba3c9-f1b3-4ea0-9839-fb0c4fb5f449"
            },
            {
              "type": "match",
              "id": "09425e95-88c5-48a8-baff-cd01751da186"
            },
            {
              "type": "match",
              "id": "75c8798e-b7ed-43d1-b56e-2786499e7c31"
            },
            {
              "type": "match",
              "id": "4692137d-6abf-4b40-8a5d-9607a88a9c81"
            },
            {
              "type": "match",
              "id": "09fcb43c-63c5-4ebe-aae1-080aee9774f3"
            },
            {
              "type": "match",
              "id": "00ddb2ad-9862-433c-9c0c-feb35cbd358f"
            },
            {
              "type": "match",
              "id": "70cbee29-1bd5-463f-9abe-45c606946490"
            },
            {
              "type": "match",
              "id": "ec6da2e8-2640-4fc3-b494-d44ea8f78f61"
            },
            {
              "type": "match",
              "id": "56c80957-bc63-4db1-b751-58540c70b864"
            },
            {
              "type": "match",
              "id": "bf47bfd9-01ef-4071-af3c-c76bdc47cae5"
            },
            {
              "type": "match",
              "id": "64ead2f9-248d-49ae-a421-2e25b76caff1"
            },
            {
              "type": "match",
              "id": "efa4fe80-9800-49b2-a6bb-11a7c0473f5a"
            },
            {
              "type": "match",
              "id": "811db2be-a280-4955-ab2b-18e784417aa7"
            },
            {
              "type": "match",
              "id": "5e50be44-15a8-4764-911f-cf9034daf0cc"
            },
            {
              "type": "match",
              "id": "cfbc2e48-e341-44d1-b19c-249b73acc336"
            },
            {
              "type": "match",
              "id": "078ac7ae-c5b7-4fab-8afe-deb6f7061d10"
            },
            {
              "type": "match",
              "id": "464e917d-6732-4956-9150-9ee2c3a830d9"
            },
            {
              "type": "match",
              "id": "6cec54d9-6324-4456-85e6-213f833361c5"
            },
            {
              "type": "match",
              "id": "645bcb9d-c00f-4f47-adce-c3a7f1f6fa81"
            },
            {
              "type": "match",
              "id": "f7585cca-715d-400b-8d51-e02969a09a24"
            },
            {
              "type": "match",
              "id": "081693cb-e471-4b8f-9eea-d9269b7f4251"
            },
            {
              "type": "match",
              "id": "e98a2b06-131a-4d60-88c7-cc29f0be9039"
            },
            {
              "type": "match",
              "id": "be5f35d5-d530-4b51-9246-8940a6fcdd52"
            },
            {
              "type": "match",
              "id": "a9de9617-20db-413f-b9c4-ab0fdd61623f"
            },
            {
              "type": "match",
              "id": "f86ff4ef-a56a-4e62-9d77-d22be6843f94"
            },
            {
              "type": "match",
              "id": "824d52d6-5ceb-4c77-a38e-2bf38c9145c6"
            },
            {
              "type": "match",
              "id": "add07017-c350-44f4-8ff0-025283b8783d"
            },
            {
              "type": "match",
              "id": "b873141f-966b-4f4e-be22-43386f8b1093"
            },
            {
              "type": "match",
              "id": "c43d8c27-9d1d-4995-a1f4-9dbac35a23b5"
            },
            {
              "type": "match",
              "id": "839b1142-6170-4574-96e4-29724cc71eaa"
            },
            {
              "type": "match",
              "id": "38d93a16-88ae-4587-a7c1-36339eb03bdd"
            },
            {
              "type": "match",
              "id": "07d576de-d787-4a66-842c-e8412b8e2731"
            },
            {
              "type": "match",
              "id": "a601a973-25eb-4f47-acea-7b3093d8a81a"
            },
            {
              "type": "match",
              "id": "327be7f6-dcb1-4d91-8401-c60b5e6fd745"
            },
            {
              "type": "match",
              "id": "b3392183-53f4-4cf0-8e3f-703ac83b4266"
            },
            {
              "type": "match",
              "id": "80bb3430-d1e4-4675-b13f-f025d1936765"
            },
            {
              "type": "match",
              "id": "076638b1-e9ba-4a3d-b0d2-c04abd221434"
            },
            {
              "type": "match",
              "id": "257b5a4c-297d-458c-9db3-27d0c61c4181"
            },
            {
              "type": "match",
              "id": "71248320-26b1-47f3-a936-8b568f99904e"
            },
            {
              "type": "match",
              "id": "d6fc292f-8c9d-48e1-a74f-fa7093d9bcfa"
            },
            {
              "type": "match",
              "id": "2919fc83-8d34-43f5-b8ce-ad63872d1741"
            },
            {
              "type": "match",
              "id": "43a5c1ef-4055-42ef-a3a8-a2abc3a2dc3c"
            },
            {
              "type": "match",
              "id": "73a49091-6e2e-44d1-abc9-7ba01d6ac9a2"
            },
            {
              "type": "match",
              "id": "32f9f592-3cbe-44b3-93b4-60a33d0d8f53"
            },
            {
              "type": "match",
              "id": "d8860043-e394-4ec7-a1b3-b7fc0823562c"
            },
            {
              "type": "match",
              "id": "e403c1aa-9f66-4587-8884-e690cddade84"
            },
            {
              "type": "match",
              "id": "17c905d9-8683-4f54-b331-b37db994685c"
            },
            {
              "type": "match",
              "id": "d4490396-7182-45a8-ad68-ae2263c942ea"
            },
            {
              "type": "match",
              "id": "e4f339ed-841b-4703-b8e7-1c4fca464483"
            },
            {
              "type": "match",
              "id": "b0cd40a3-8c97-44ef-aa24-bfc7320a77ee"
            },
            {
              "type": "match",
              "id": "4ed4677b-8e27-41e9-9766-48fa49d844df"
            },
            {
              "type": "match",
              "id": "b5174bc9-23be-4ad2-88a7-c541869075c8"
            },
            {
              "type": "match",
              "id": "451e5aac-082a-4462-9f22-688797a5c0e2"
            },
            {
              "type": "match",
              "id": "23f73cb0-ca3c-4d86-bdfe-e1ffb5993022"
            },
            {
              "type": "match",
              "id": "58877e97-290c-4bb0-83c6-6d7d28859437"
            },
            {
              "type": "match",
              "id": "c07973e8-e43b-45c9-9097-3cdbdf68fceb"
            },
            {
              "type": "match",
              "id": "1f4cfce6-f052-4f91-abc8-1063e8f163bf"
            },
            {
              "type": "match",
              "id": "13b7212a-e933-4a13-ac61-6c9aaf2937ee"
            },
            {
              "type": "match",
              "id": "c4867095-8105-4a3e-ab35-31c40851d045"
            },
            {
              "type": "match",
              "id": "cbc69aec-74f2-4008-a4cc-b30d516c8098"
            },
            {
              "type": "match",
              "id": "b1477fdd-b85a-4984-b4ae-8005715baf12"
            },
            {
              "type": "match",
              "id": "89a2acb1-40c8-4782-8021-fa4a8fc8c4be"
            },
            {
              "type": "match",
              "id": "4da0e535-75d0-42ee-86af-104d63a479ec"
            },
            {
              "type": "match",
              "id": "5918d4f5-ee70-4077-bbbd-9d0d22155e55"
            },
            {
              "type": "match",
              "id": "6b26dd1c-5ef1-4fed-85f0-91213effbd7b"
            },
            {
              "type": "match",
              "id": "6c92eca2-2280-4945-8412-f201cf05c3ea"
            },
            {
              "type": "match",
              "id": "e65a8f08-7a54-41fa-9b65-3b72dc782d16"
            },
            {
              "type": "match",
              "id": "955a1d84-2532-4104-9527-7f101ee9f86c"
            }
          ]
        }
      },
      "links": {
        "self": "https://api.playbattlegrounds.com/shards/steam/players/account.c0e530e9b7244b358def282782f893af",
        "schema": ""
      }
    }
  ],
  "links": {
    "self": "https://api.pubg.com/shards/steam/players?filter[playerNames]=WackyJacky101"
  },
  "meta": {}
}
`
func TestParsePlayer(t *testing.T) {
	players, err := ParsePlayers(testPlayerStr); if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, players, 1)
	assert.Equal(t, "account.c0e530e9b7244b358def282782f893af", players[0].ID)
	assert.Equal(t, "WackyJacky101", players[0].Name)
	assert.Equal(t, "steam", players[0].ShardID)
	assert.Equal(t, "bluehole-pubg", players[0].TitleID)
	assert.Equal(t, "4b1ea7d9-b512-482b-a025-ae6ee5403c64", players[0].Matches[0].ID)
	assert.Equal(t, "2020-07-01 14:38:28 +0000 UTC", players[0].CreatedAt.String())
}
