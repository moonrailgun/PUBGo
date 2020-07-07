package schema

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testSamplesStr = `
{
  "data": {
    "type": "sample",
    "id": "2da5b82b-72ed-406f-9216-7c382a672f4d",
    "attributes": {
      "createdAt": "2020-07-06T00:00:00Z",
      "titleId": "bluehole-pubg",
      "shardId": "steam"
    },
    "relationships": {
      "matches": {
        "data": [
          {
            "type": "match",
            "id": "b8b0003d-9722-4aa9-ab88-7bdf19c65a3e"
          },
          {
            "type": "match",
            "id": "61d9feae-8b56-4302-bbe9-333c495e1cff"
          },
          {
            "type": "match",
            "id": "41f6034d-2d61-4806-95b5-a1feb48645ce"
          },
          {
            "type": "match",
            "id": "1bf66bed-086e-419d-a29a-dc2ce75fa549"
          },
          {
            "type": "match",
            "id": "a1caad91-8ea6-499c-af6a-cc00613f3eec"
          },
          {
            "type": "match",
            "id": "5488aabe-d405-45b1-a755-cc245d05e969"
          },
          {
            "type": "match",
            "id": "80296700-2bef-4a99-b2e2-571a6fa798a9"
          },
          {
            "type": "match",
            "id": "27b744a9-3bec-4f44-8072-8da6bab649f2"
          },
          {
            "type": "match",
            "id": "1480b1aa-1887-45bf-9b89-8c04b57cd3d0"
          },
          {
            "type": "match",
            "id": "d267efb8-7e68-4aaf-8c4a-709ca1dd09ab"
          },
          {
            "type": "match",
            "id": "2d68deb1-8625-45ff-83ae-95f6c8e3adaa"
          },
          {
            "type": "match",
            "id": "4beceb10-b40e-47a8-98b9-9edcee74563b"
          },
          {
            "type": "match",
            "id": "8edaadf4-9914-40ac-aa06-704a4aeeb8be"
          },
          {
            "type": "match",
            "id": "d108dfcf-dc5c-4438-96e6-388ff1851dd1"
          },
          {
            "type": "match",
            "id": "4cf5f726-d8bf-40cd-b33d-29b979207821"
          },
          {
            "type": "match",
            "id": "68fa6ba5-73fe-43f7-9070-ea43b005e4d5"
          },
          {
            "type": "match",
            "id": "f81edb05-ecb6-45a9-8642-0c18d48e3499"
          },
          {
            "type": "match",
            "id": "14e609df-ae37-4611-9985-3ca53394b54d"
          },
          {
            "type": "match",
            "id": "a071c5b5-90d7-459d-8a33-45d0e098ff24"
          },
          {
            "type": "match",
            "id": "bae6a48f-000f-45e7-8b27-ce0cb869e681"
          },
          {
            "type": "match",
            "id": "a43eb0d9-0cd0-4907-acb5-95ee2d18c551"
          },
          {
            "type": "match",
            "id": "82f0e582-7364-4e39-b434-9418ad493ce2"
          },
          {
            "type": "match",
            "id": "c87ddb32-02f5-4c1f-bbcb-1a17769a89ea"
          },
          {
            "type": "match",
            "id": "ace37f0e-e65f-44c5-8db7-1f382e1b3ae8"
          },
          {
            "type": "match",
            "id": "cb349e45-58d0-4bc1-b0db-ff52b47f8a9c"
          },
          {
            "type": "match",
            "id": "470790e1-3e1b-4f3e-a154-f143ce11124a"
          },
          {
            "type": "match",
            "id": "0d84e063-f5e9-4458-b1d4-1c8a6c67d253"
          },
          {
            "type": "match",
            "id": "ec38cf36-88fb-4669-b257-9f12cbdb2a03"
          },
          {
            "type": "match",
            "id": "858f5178-d976-4436-b4d8-f8229b501ddb"
          },
          {
            "type": "match",
            "id": "d8093b36-b4b4-43ae-945c-6860563e8967"
          },
          {
            "type": "match",
            "id": "478bcac6-9a45-454f-bb38-63a723f9bdba"
          },
          {
            "type": "match",
            "id": "99aa550d-7579-48c6-92e2-efda21623006"
          },
          {
            "type": "match",
            "id": "1b90076f-f567-413f-b205-76bb9db34626"
          },
          {
            "type": "match",
            "id": "db0444db-7f73-402a-be11-2f363f67f6bc"
          },
          {
            "type": "match",
            "id": "b94db6c6-e9e3-49e0-825c-0942239114b6"
          },
          {
            "type": "match",
            "id": "80b82659-65a6-4ba9-a298-a94cab615cf3"
          },
          {
            "type": "match",
            "id": "56af789f-fc4d-4c49-afde-7f86805ff755"
          },
          {
            "type": "match",
            "id": "d7d5e5c5-0096-4848-8dca-8251bca06e4b"
          },
          {
            "type": "match",
            "id": "4c0f2c6e-cf56-45f2-9df3-dfd6e37cae46"
          },
          {
            "type": "match",
            "id": "271dd345-5032-4afc-a961-89ca9ce48353"
          },
          {
            "type": "match",
            "id": "cf597238-2e7c-49f1-bc7e-4495bd3cb34b"
          },
          {
            "type": "match",
            "id": "3fb05561-4f42-416b-a07a-ab416d761a5f"
          },
          {
            "type": "match",
            "id": "f4d6a713-2afb-407a-a253-2d23738f2096"
          },
          {
            "type": "match",
            "id": "7010ce37-a520-40d9-92e0-3308c54705cd"
          },
          {
            "type": "match",
            "id": "14324380-bd91-449a-bb98-845876e3d85b"
          },
          {
            "type": "match",
            "id": "70139e6a-0300-42d0-9af0-2f8f5dfeed26"
          },
          {
            "type": "match",
            "id": "7f70e588-a693-40b6-bda5-5fd5250e5cf3"
          },
          {
            "type": "match",
            "id": "0839a650-71a6-44da-988e-9c0443a9a69a"
          },
          {
            "type": "match",
            "id": "052f113e-36b0-4df9-b888-1b69e68e252a"
          },
          {
            "type": "match",
            "id": "3809d0ce-7639-44df-8837-b1db85df4bf0"
          },
          {
            "type": "match",
            "id": "bf0e308c-8e2f-4e7f-96ee-809c98e6524c"
          },
          {
            "type": "match",
            "id": "f1fea45f-3be3-4be4-af4b-6d55d28b0767"
          },
          {
            "type": "match",
            "id": "9601df0e-23a3-4346-b645-ed48c5ae0ab0"
          },
          {
            "type": "match",
            "id": "a8c8cb46-0d7a-448b-8e18-0fac761eade6"
          },
          {
            "type": "match",
            "id": "5530ca21-a639-43b0-93c2-0bba90be3ec3"
          },
          {
            "type": "match",
            "id": "05c379a8-4a2c-4690-a124-6a0134331c18"
          },
          {
            "type": "match",
            "id": "778e5673-e7ff-4974-9716-641657c59fac"
          },
          {
            "type": "match",
            "id": "f07efe06-40d5-4ebf-ae4e-c1c5857b0ebd"
          },
          {
            "type": "match",
            "id": "881eb6d7-fc6d-4fb5-9989-a5770963a43c"
          },
          {
            "type": "match",
            "id": "47d8f0d4-7cdd-4637-906e-c639774f4a7c"
          },
          {
            "type": "match",
            "id": "3c86bb47-1c80-41fe-b0ef-c243810b47e0"
          },
          {
            "type": "match",
            "id": "2fb54768-d712-466f-afce-8edc9abfbe11"
          },
          {
            "type": "match",
            "id": "4dada9d5-4545-492b-b0de-24e2cc119922"
          },
          {
            "type": "match",
            "id": "138562a9-04f1-4e57-a5e7-139ec187ca29"
          },
          {
            "type": "match",
            "id": "c112467d-2816-4f6c-b0b1-ddf95d88793f"
          },
          {
            "type": "match",
            "id": "0a1aafc6-6a7d-4539-8ca9-b54d970b9f12"
          },
          {
            "type": "match",
            "id": "4446650e-088d-4ee4-90c5-da951a787ff6"
          },
          {
            "type": "match",
            "id": "d0ffc1e8-5386-4c3d-a49e-4757f81ac75b"
          },
          {
            "type": "match",
            "id": "9fe27e63-0af0-40be-83a0-176c021cceac"
          },
          {
            "type": "match",
            "id": "46d28284-8449-454e-a231-30a62437d531"
          },
          {
            "type": "match",
            "id": "928da887-7ea5-46e7-bbee-867a8b46ad52"
          },
          {
            "type": "match",
            "id": "b011ba96-5647-489a-962e-be71e5259def"
          },
          {
            "type": "match",
            "id": "11412e32-1655-4131-a229-097fd3a10804"
          },
          {
            "type": "match",
            "id": "3e864e80-28b6-48cf-ab55-a601e1bb032c"
          },
          {
            "type": "match",
            "id": "bb3fa40a-8716-4ffe-b37f-62dfd9518a0f"
          },
          {
            "type": "match",
            "id": "1e3684b7-e129-4c54-abba-f349efdbfa43"
          },
          {
            "type": "match",
            "id": "19b0adff-5ef1-4ad4-883d-7deb12e33a28"
          },
          {
            "type": "match",
            "id": "1bd81c85-601c-4fb2-8782-816e81a60a06"
          },
          {
            "type": "match",
            "id": "c94be071-310d-474b-b2d8-47c86b04113d"
          },
          {
            "type": "match",
            "id": "7132588a-c49d-4047-896c-35215aa41926"
          },
          {
            "type": "match",
            "id": "aa01c594-601d-4269-835e-4ceefacfae97"
          },
          {
            "type": "match",
            "id": "453ac56c-4b59-4743-89d7-b1a7822ce7c9"
          },
          {
            "type": "match",
            "id": "7c496bef-e0bd-423f-a6fd-25c8cb3a1d6e"
          },
          {
            "type": "match",
            "id": "7269ef8f-4681-4e07-9133-c3ca19c2dc8e"
          },
          {
            "type": "match",
            "id": "2c5333bf-23bd-41d8-b17d-10d6cdd61c0e"
          },
          {
            "type": "match",
            "id": "86e42113-4033-4962-b15b-3b4c3a9e3bfe"
          },
          {
            "type": "match",
            "id": "ea5419dd-b185-4767-a8d3-cd8e43eb9b9f"
          },
          {
            "type": "match",
            "id": "38f87e52-3233-40c2-91c5-89e10b0996d7"
          },
          {
            "type": "match",
            "id": "e3b1ac60-214f-4a00-b1b6-bfc503ce612c"
          },
          {
            "type": "match",
            "id": "3681f978-5398-4807-b733-85a7e5fbb444"
          },
          {
            "type": "match",
            "id": "41eec2df-9f2b-4f41-a260-b0d195fe48f8"
          },
          {
            "type": "match",
            "id": "2c48e2f5-c58f-4212-80a8-cf6cd429d97a"
          },
          {
            "type": "match",
            "id": "e698416c-b1df-435b-8e05-18cfd09786f6"
          },
          {
            "type": "match",
            "id": "616e5aa7-9b20-4368-96d7-6ff9411bfc83"
          },
          {
            "type": "match",
            "id": "6fc7fe8c-4b26-45eb-8c6f-52393de987ac"
          },
          {
            "type": "match",
            "id": "b3aee955-a854-43e7-b5c0-4d00a7402ae3"
          },
          {
            "type": "match",
            "id": "73728d7e-d52d-4f64-91bc-6aad377b7b57"
          },
          {
            "type": "match",
            "id": "7ca995ef-3bb8-476d-a52e-55aa650bb006"
          },
          {
            "type": "match",
            "id": "bfea86cd-ac2a-4526-b6f4-5b92f00d92e5"
          },
          {
            "type": "match",
            "id": "7a90874d-4379-4ba7-944d-98fd12959eef"
          },
          {
            "type": "match",
            "id": "1125d5ae-0c9c-4d16-9c7a-deafb0033e12"
          },
          {
            "type": "match",
            "id": "2a0dff1b-dc42-406e-9c69-9739289ce33a"
          },
          {
            "type": "match",
            "id": "21f2c48b-cd42-4fe6-8467-cd0cf57f23ca"
          },
          {
            "type": "match",
            "id": "b2408e78-f7bf-45e9-9f05-d6b870a65011"
          },
          {
            "type": "match",
            "id": "5a673dd8-df6c-4678-80a5-f6489b3ea0fc"
          },
          {
            "type": "match",
            "id": "967d7ee2-a0ad-45c6-a036-25b26e7c41e5"
          },
          {
            "type": "match",
            "id": "4dfa9229-0be5-4a46-b485-c875ea411d23"
          },
          {
            "type": "match",
            "id": "0c92ad7b-03b6-4df9-990c-fdf44c5c7396"
          },
          {
            "type": "match",
            "id": "2ec8c474-6c70-4856-800f-05b4a97c7486"
          },
          {
            "type": "match",
            "id": "a1635a6e-ca96-48a6-a6b6-1644141bff21"
          },
          {
            "type": "match",
            "id": "f2e0d0bf-175c-461f-8b60-327f80ff6152"
          },
          {
            "type": "match",
            "id": "85586f68-d3d2-442e-a39c-051b9a45630d"
          },
          {
            "type": "match",
            "id": "0423d698-ec3c-4ddb-9a3d-6dc30f477477"
          },
          {
            "type": "match",
            "id": "5fd32723-b8bc-473f-83ef-acb8eaa65726"
          },
          {
            "type": "match",
            "id": "5b219f1e-5b00-42cd-bcaa-b4c5164cd9eb"
          },
          {
            "type": "match",
            "id": "6219ff4d-63fe-475d-b0c8-24b6d705803a"
          },
          {
            "type": "match",
            "id": "e8c33baa-4224-456a-8ec3-1e34d6c7937c"
          },
          {
            "type": "match",
            "id": "b87c8134-a67c-4fe3-8a93-46dc90dd5f30"
          },
          {
            "type": "match",
            "id": "f156d102-5f8a-4abc-aa4a-4f15eb790c6a"
          },
          {
            "type": "match",
            "id": "b9e6c9dd-6ba3-4eeb-bc66-53cea0382d3e"
          },
          {
            "type": "match",
            "id": "db02dead-e611-4795-9bed-007cf0d1fd39"
          },
          {
            "type": "match",
            "id": "5665eadb-d78a-4c27-a570-3206055bf618"
          },
          {
            "type": "match",
            "id": "cd97ef53-e747-41c8-a753-4bd840da8bc2"
          },
          {
            "type": "match",
            "id": "fe0a2f1a-4a07-4868-ae97-d9482cf55e9d"
          },
          {
            "type": "match",
            "id": "172d875c-4409-44e1-a2f2-5dfba97f4267"
          },
          {
            "type": "match",
            "id": "42584793-943d-4e47-81c4-008e40f14c72"
          },
          {
            "type": "match",
            "id": "76906d5a-3cbd-4e41-be3f-876a30db32f9"
          },
          {
            "type": "match",
            "id": "e1bf93a3-d8f9-4374-ab64-2c7f73af06fa"
          },
          {
            "type": "match",
            "id": "15da3838-ff1c-46bc-bd05-4f61d00295c9"
          },
          {
            "type": "match",
            "id": "11f51b05-7dd6-4229-ae7f-0394db70b48b"
          },
          {
            "type": "match",
            "id": "327e27e1-213f-46be-a12d-efa6d89f12c5"
          },
          {
            "type": "match",
            "id": "9cb61611-2405-46e8-b94a-344625a9f15c"
          },
          {
            "type": "match",
            "id": "dc84e701-d651-4e90-ac2b-3fd16c587fed"
          },
          {
            "type": "match",
            "id": "12ee367b-8637-4062-a7ae-926bdda11a74"
          },
          {
            "type": "match",
            "id": "b49eaa1e-654b-47a0-9764-efcaeaf23f2f"
          },
          {
            "type": "match",
            "id": "d5106bb6-dafa-4490-ba6e-15c2cb05d18a"
          },
          {
            "type": "match",
            "id": "9ac4cef2-d6c0-48eb-90e6-bdb99fdfd055"
          },
          {
            "type": "match",
            "id": "9cb705fc-36aa-4cb3-a63b-038c8544bc42"
          },
          {
            "type": "match",
            "id": "74958768-907c-4787-b062-3208aaa1d007"
          },
          {
            "type": "match",
            "id": "e1e81ebf-49c4-4c37-9764-b23bb350dd15"
          },
          {
            "type": "match",
            "id": "9c532d9c-ac02-4de3-be79-73b7d8dd0f26"
          },
          {
            "type": "match",
            "id": "0e428689-97d8-457e-9b36-f0d82a2e994a"
          },
          {
            "type": "match",
            "id": "05eac0c7-bf6d-42c8-9c1d-cda3cceb2751"
          },
          {
            "type": "match",
            "id": "1d43b81b-a6b0-4193-aa32-d7beef1be8b6"
          },
          {
            "type": "match",
            "id": "b3765a75-df36-407d-8a96-e06b095948b9"
          },
          {
            "type": "match",
            "id": "b8167011-6c97-4c24-9684-035dd7468636"
          },
          {
            "type": "match",
            "id": "f67f7536-1b66-42f9-b306-d9153647211b"
          },
          {
            "type": "match",
            "id": "9b4d964a-1334-4fea-8617-f7361a96ec47"
          },
          {
            "type": "match",
            "id": "06d31f89-d365-40e9-86fd-82dfa8dad06b"
          },
          {
            "type": "match",
            "id": "cc1fe688-9a76-401c-b66f-fe42499fa6a7"
          },
          {
            "type": "match",
            "id": "45b08b99-d0e6-4855-ba5d-cc6137787cd0"
          },
          {
            "type": "match",
            "id": "6e1e7f9a-9928-43f8-b8ab-97a2d75a22d9"
          },
          {
            "type": "match",
            "id": "03a19bf5-6dc4-4f70-9e03-3b8476de53ba"
          },
          {
            "type": "match",
            "id": "ae50704f-1828-4c54-8337-1b15d9d2d218"
          },
          {
            "type": "match",
            "id": "41e67339-2d79-46da-90bd-7b5808034483"
          },
          {
            "type": "match",
            "id": "f1de5a33-2188-4812-8944-2747da0afe79"
          },
          {
            "type": "match",
            "id": "2cf49ad9-3891-4fb2-8408-a9e2cd26f6d4"
          },
          {
            "type": "match",
            "id": "a1fe2ecd-0666-44b9-ab34-54f8c0f39c83"
          },
          {
            "type": "match",
            "id": "ecbdf016-3c00-4376-9ebb-f16184625a61"
          },
          {
            "type": "match",
            "id": "3eaf4ad4-e1e6-4f18-89b2-d039c7ff9252"
          },
          {
            "type": "match",
            "id": "63944fcb-2cad-4b54-9b51-75856d8e5b83"
          },
          {
            "type": "match",
            "id": "d4aa784e-eea5-435e-950c-34e790a41dc7"
          },
          {
            "type": "match",
            "id": "41854edf-a6ed-4b87-9a6a-1b05d4eeb225"
          },
          {
            "type": "match",
            "id": "a8ca3334-f3c6-4737-a028-ffca98ec5387"
          },
          {
            "type": "match",
            "id": "c89928dd-fd90-47da-ab3c-f2451a7f12e6"
          },
          {
            "type": "match",
            "id": "3f70898e-a9c0-4f52-9a5f-369696e433a5"
          },
          {
            "type": "match",
            "id": "478e96e3-0f60-42fe-b69a-4e81ef97fcad"
          },
          {
            "type": "match",
            "id": "38284732-8ba9-49d4-99de-7eafaef1cd4d"
          },
          {
            "type": "match",
            "id": "f3af11f6-fb4d-45ae-9304-9848b54d45a9"
          },
          {
            "type": "match",
            "id": "7adbb32e-5c5f-48a4-9877-4764a69d7c25"
          },
          {
            "type": "match",
            "id": "e057a6e9-6165-46ef-917f-ff564bc487b6"
          },
          {
            "type": "match",
            "id": "03f96c28-d251-4f17-8151-e10e03f69857"
          },
          {
            "type": "match",
            "id": "00662f21-7e07-4a74-aa17-bd449c21b964"
          },
          {
            "type": "match",
            "id": "541091d3-482b-421e-a24e-61b74e540767"
          },
          {
            "type": "match",
            "id": "475730d3-55d9-487e-9d8e-e36cf1e169fc"
          },
          {
            "type": "match",
            "id": "02a587ec-3ff7-4a10-84b8-fa32d06ce835"
          },
          {
            "type": "match",
            "id": "648eec7d-03b0-4896-823c-1e5982d24197"
          },
          {
            "type": "match",
            "id": "ac20f462-3699-4741-9f55-25bb041e4d5a"
          },
          {
            "type": "match",
            "id": "735eccbe-a8ac-4ee9-975d-fb017e6c129c"
          },
          {
            "type": "match",
            "id": "6cec879c-0e07-41e1-9fe6-72ee4e4a65a9"
          },
          {
            "type": "match",
            "id": "10e27d94-8365-45d5-817f-a2725f9ab289"
          },
          {
            "type": "match",
            "id": "0658544b-7a2d-4f7b-b471-794688cc32d1"
          },
          {
            "type": "match",
            "id": "cbdbf63b-7bcb-4ccf-94c6-76c923b29ed7"
          },
          {
            "type": "match",
            "id": "dde8633f-e81a-454e-aaf1-81b3d2b10b79"
          },
          {
            "type": "match",
            "id": "7a4b0c99-b637-4900-8a79-7dab6975a925"
          },
          {
            "type": "match",
            "id": "f41e87d1-930d-485b-988e-a8a0e6289849"
          },
          {
            "type": "match",
            "id": "e80be569-b255-4006-bb62-2b1140be8236"
          },
          {
            "type": "match",
            "id": "fe3dd923-5690-4d7e-9ff3-f165b3601ccf"
          },
          {
            "type": "match",
            "id": "f89f2553-2828-4b9a-91a8-77e1fbfcc11e"
          },
          {
            "type": "match",
            "id": "1edc5a15-12f8-4ca2-970b-0625b46e0f0b"
          },
          {
            "type": "match",
            "id": "1bf3d1c2-399b-447d-a73e-18ad3bdec9a5"
          },
          {
            "type": "match",
            "id": "154e8676-e2ac-480d-a72c-914b99162daf"
          },
          {
            "type": "match",
            "id": "028a3607-0a7c-4571-b5e0-693c6ed50799"
          },
          {
            "type": "match",
            "id": "bc9c74c6-1db0-4ea8-9fd2-b195cac5f1d5"
          },
          {
            "type": "match",
            "id": "587567ee-a0a4-4cdf-80e7-f811383a2347"
          },
          {
            "type": "match",
            "id": "6323784b-85c9-4b0a-a3d7-0fe78b44b9b3"
          },
          {
            "type": "match",
            "id": "86bfe979-c7af-478a-b295-d4783cb74336"
          },
          {
            "type": "match",
            "id": "f322e005-5805-4e89-99c2-9795fe589217"
          },
          {
            "type": "match",
            "id": "82008cd9-4e6b-43b7-b872-d6a72cb7b907"
          },
          {
            "type": "match",
            "id": "8d1f5444-c36c-4b38-bc91-4d1ff99937d4"
          },
          {
            "type": "match",
            "id": "a3baa4cf-2b56-48cb-8969-a800b53a5533"
          },
          {
            "type": "match",
            "id": "13227190-02b6-475e-916b-b84257d7a7a9"
          },
          {
            "type": "match",
            "id": "4e7f85b0-baab-4201-be68-c1729c18136f"
          },
          {
            "type": "match",
            "id": "950f90d6-b3e5-4ef5-957b-8cbfed371788"
          },
          {
            "type": "match",
            "id": "4860ea25-9977-4677-9046-b15535d66956"
          },
          {
            "type": "match",
            "id": "d75f08e7-c31e-419f-868a-088224f94caf"
          },
          {
            "type": "match",
            "id": "6d0c93f1-0b74-4970-80e5-458a7f3b0730"
          },
          {
            "type": "match",
            "id": "7d317364-e2f4-4592-a0ad-af8f71c41e8d"
          },
          {
            "type": "match",
            "id": "2db28684-aa0c-4ad5-b07d-cb496594e113"
          },
          {
            "type": "match",
            "id": "19e04ea1-c8ce-4697-9da7-b1402f978e7f"
          },
          {
            "type": "match",
            "id": "ba873a30-2edc-44ba-ac2e-f6c04cecb8f9"
          },
          {
            "type": "match",
            "id": "ba3b6457-e566-46b0-b6a4-f7f7fe365c3a"
          },
          {
            "type": "match",
            "id": "725fbd8c-6e94-43b4-a838-044aab5e54df"
          },
          {
            "type": "match",
            "id": "c2a6456a-6010-4e9b-ac14-cdf1f09d4255"
          },
          {
            "type": "match",
            "id": "5fde456f-197d-4b4d-b6e4-0aeb8de0b91c"
          },
          {
            "type": "match",
            "id": "7ea0ee00-2f28-4787-9e87-a00323ac40cb"
          },
          {
            "type": "match",
            "id": "dbf99db0-8070-498f-83d5-37f5c440e736"
          },
          {
            "type": "match",
            "id": "80781c8c-f1f3-4785-9757-bc8c01ccb652"
          },
          {
            "type": "match",
            "id": "290fff57-d0c0-41bb-8308-aa9cdf348e1c"
          },
          {
            "type": "match",
            "id": "6d710265-40a8-4956-a0c3-c35095973666"
          },
          {
            "type": "match",
            "id": "13dbaff5-cf9b-4959-bd16-7f0f9899c4a8"
          },
          {
            "type": "match",
            "id": "b3fe21f0-ddf0-401f-b2a3-66e13011054a"
          },
          {
            "type": "match",
            "id": "f83e139d-0705-4467-9fac-73a2e58e650a"
          },
          {
            "type": "match",
            "id": "36847f1e-8ede-40f2-a7d7-bfbf5fb6c09e"
          },
          {
            "type": "match",
            "id": "a0592c45-c433-40da-84d9-4985343f13ae"
          },
          {
            "type": "match",
            "id": "c497a889-fce9-4dfd-a622-f3911638740e"
          },
          {
            "type": "match",
            "id": "98a26fe7-71d6-40f6-bafd-4bc1d6f1bfdc"
          },
          {
            "type": "match",
            "id": "5011f0f6-b3ce-4823-81b9-a548926c2d0e"
          },
          {
            "type": "match",
            "id": "26497e76-12bf-4818-84df-19ce062d9571"
          },
          {
            "type": "match",
            "id": "e2e5ded9-3ff6-4f02-8acf-41fbe8847ef7"
          },
          {
            "type": "match",
            "id": "911ff879-53e4-435b-89fb-c2d3ac838ba6"
          },
          {
            "type": "match",
            "id": "81a7c652-6b76-4ab4-bec8-c556f21d50ca"
          },
          {
            "type": "match",
            "id": "4ad13af6-1622-4b29-b051-c2e9480d0990"
          },
          {
            "type": "match",
            "id": "3448af3c-10f3-4143-afc5-fcbdd95b49ff"
          },
          {
            "type": "match",
            "id": "33a68420-4cdc-47d0-90a8-f97e5a81ccfd"
          },
          {
            "type": "match",
            "id": "f993debc-831e-48b9-97b1-af511ba87f80"
          },
          {
            "type": "match",
            "id": "cbbe5d08-3e47-47ab-99f1-f5b7225ea6e6"
          },
          {
            "type": "match",
            "id": "00769a39-b6f2-47d4-8181-4b3116ea7ddf"
          },
          {
            "type": "match",
            "id": "58161bba-6796-4e34-a42c-7515d48035be"
          },
          {
            "type": "match",
            "id": "a3502a0b-3de5-4fd4-9ec2-b91c587e83f3"
          },
          {
            "type": "match",
            "id": "735a9582-1767-48ff-b746-87c5ea0b5aaa"
          },
          {
            "type": "match",
            "id": "94b7127c-64f6-44eb-a98e-de545156b6b9"
          },
          {
            "type": "match",
            "id": "59cb1ecf-e332-4e77-8e1c-41275e52afc3"
          },
          {
            "type": "match",
            "id": "2ac41b26-6480-49c6-ac73-b1cceb343df1"
          },
          {
            "type": "match",
            "id": "643be721-642c-48ce-8994-198516860221"
          },
          {
            "type": "match",
            "id": "b120dd1a-a99c-4393-ad7f-e3fe100aa6ac"
          },
          {
            "type": "match",
            "id": "8d173d27-275d-45ed-b4b6-15ac958cc2d5"
          },
          {
            "type": "match",
            "id": "8b8faa9f-709a-40c5-bb2b-5959835e5f48"
          },
          {
            "type": "match",
            "id": "5f0d8f05-b22e-4f2b-a99d-0f02c63afede"
          },
          {
            "type": "match",
            "id": "e0906394-f96d-4113-8f31-a04abe4b5309"
          },
          {
            "type": "match",
            "id": "54e4bb56-e9f7-4790-8582-8c40cbf1d121"
          },
          {
            "type": "match",
            "id": "33686941-6be2-4d3f-80c8-b73fb3d7e893"
          },
          {
            "type": "match",
            "id": "4e2806e7-48ef-4ddb-b430-6633b801d940"
          },
          {
            "type": "match",
            "id": "c3250ab9-8bbf-45f2-9b00-16ada9a24a1d"
          },
          {
            "type": "match",
            "id": "115ec68b-fac1-4cfb-aceb-db013543aa33"
          },
          {
            "type": "match",
            "id": "0f6e0fea-93fc-41aa-aee3-f8eebac2eaad"
          },
          {
            "type": "match",
            "id": "288056e7-be7a-476e-be08-14a595c4ded6"
          },
          {
            "type": "match",
            "id": "6b3b6963-1b3f-44a5-8ccb-69d19eebad6a"
          },
          {
            "type": "match",
            "id": "0cefef6f-113f-418b-852c-d820818039bb"
          },
          {
            "type": "match",
            "id": "6c49beb6-6f4f-47aa-bc44-e4008fc329c9"
          },
          {
            "type": "match",
            "id": "80ffc7b6-4d07-4ceb-96a5-b3f89c621cc0"
          },
          {
            "type": "match",
            "id": "1dc6e165-7635-45ca-af97-6d633ce594dc"
          },
          {
            "type": "match",
            "id": "d09d23e5-f494-47aa-b641-8e6f8547c2a1"
          },
          {
            "type": "match",
            "id": "29b04066-f3f1-4d32-97a7-3a9efd12a943"
          },
          {
            "type": "match",
            "id": "d280d280-76ff-4873-800c-a1d43a7ecea5"
          },
          {
            "type": "match",
            "id": "702efd30-583d-407b-9db1-3cfc51bb29d3"
          },
          {
            "type": "match",
            "id": "045ea50a-f64f-4089-a4e6-3fab01984806"
          },
          {
            "type": "match",
            "id": "2eec8b91-dea4-4de6-95f0-a71736912614"
          },
          {
            "type": "match",
            "id": "c379d66b-c5b5-4b4d-9e91-de4545ddbcb0"
          },
          {
            "type": "match",
            "id": "a4a95b74-2034-4f6e-ad6f-5654370483b9"
          },
          {
            "type": "match",
            "id": "988ddddb-adf9-4d78-8580-4d19c58bf645"
          },
          {
            "type": "match",
            "id": "b2448ae6-1d3e-43f9-9627-3682e130da10"
          },
          {
            "type": "match",
            "id": "ee6663c1-a06d-498e-814d-5854ba0ef06b"
          },
          {
            "type": "match",
            "id": "e0b3f548-af26-47ea-a87f-07fecbe239ce"
          },
          {
            "type": "match",
            "id": "fa0daee3-3f28-4534-9cf4-f4a60e002edd"
          },
          {
            "type": "match",
            "id": "230a0245-cd7d-44ce-bf8c-cc168bfbccdd"
          },
          {
            "type": "match",
            "id": "89927345-d6bf-45a1-8c24-ff18b3bd908f"
          },
          {
            "type": "match",
            "id": "fef43d2b-eff2-402d-8b65-6df654f45c60"
          },
          {
            "type": "match",
            "id": "62623377-b3f0-41cd-a5d4-2b6b7e4b109e"
          },
          {
            "type": "match",
            "id": "ebc8be80-d3d3-4ec9-ab51-7c4f3bf1bf9a"
          },
          {
            "type": "match",
            "id": "ecf40556-e85c-4c19-88cc-d5b68b083cfe"
          },
          {
            "type": "match",
            "id": "e01bfc9c-47b7-4964-8632-2bf9fa05bbbc"
          },
          {
            "type": "match",
            "id": "cb50467f-9a7f-4392-be05-135d3df489c5"
          },
          {
            "type": "match",
            "id": "1ecf859e-cae6-437c-9d3d-a49ad19dd752"
          },
          {
            "type": "match",
            "id": "c5933026-5dae-4b2a-a008-48340a160bdf"
          },
          {
            "type": "match",
            "id": "59baf1dc-db45-4fbf-9862-35f14bd51c7d"
          },
          {
            "type": "match",
            "id": "38f98883-c147-4d83-aa7c-2642cbc4a6a2"
          },
          {
            "type": "match",
            "id": "e9081569-c5d0-4602-85c8-cb7679e92531"
          },
          {
            "type": "match",
            "id": "d4778b36-eb86-4e94-96f7-2a80dc499c77"
          },
          {
            "type": "match",
            "id": "5def84fb-969b-4114-884a-2f3333b7c0bc"
          },
          {
            "type": "match",
            "id": "08e17474-2aed-48c8-b499-0eab7d25ed66"
          },
          {
            "type": "match",
            "id": "8b3715cb-7879-4760-86ff-608919aa3883"
          },
          {
            "type": "match",
            "id": "e4519e25-e38f-43f0-8287-f727806d2903"
          },
          {
            "type": "match",
            "id": "8617a29f-35ee-4810-b06a-7a7edfc33a62"
          },
          {
            "type": "match",
            "id": "f5a7016d-5663-4f7e-908e-daf5990b837c"
          },
          {
            "type": "match",
            "id": "be3d2d4e-54ad-443e-bf57-daf8219710b3"
          },
          {
            "type": "match",
            "id": "d2d6fe50-95bf-42bf-b40f-44fabc8e893d"
          },
          {
            "type": "match",
            "id": "23eca41b-ca8c-49ce-8310-760a2021b1a7"
          },
          {
            "type": "match",
            "id": "4596b606-c477-4cdb-abda-b80e6c17524a"
          },
          {
            "type": "match",
            "id": "a2504384-d2ab-4ee4-9550-5b3d0ce4cc4c"
          },
          {
            "type": "match",
            "id": "c4b08c6a-f412-4cac-ae76-9df5fe967380"
          },
          {
            "type": "match",
            "id": "03a84bc5-ec90-403a-8ee0-06e6886a0427"
          },
          {
            "type": "match",
            "id": "3bdc6d36-3f34-4b94-8260-4c99dbbefb90"
          },
          {
            "type": "match",
            "id": "9a8a49bf-1af9-4ee8-a59f-ffb26b66844e"
          },
          {
            "type": "match",
            "id": "f1916fb0-adb8-4b14-888a-4ffabdfe7d5e"
          },
          {
            "type": "match",
            "id": "41022995-698d-4fb0-bac7-6d60fa0eba96"
          },
          {
            "type": "match",
            "id": "ea0dd644-ae68-44a5-af1f-35f0edf0f4d4"
          },
          {
            "type": "match",
            "id": "ff1d2bf1-edff-4a2e-bc5f-ac059b717f84"
          },
          {
            "type": "match",
            "id": "7a26599b-f622-4851-a9fd-73d2530438fd"
          },
          {
            "type": "match",
            "id": "023c8e10-a338-4d82-ac1e-d0f875f44605"
          },
          {
            "type": "match",
            "id": "a33d0862-ad06-4e65-8cca-29baf3205f90"
          },
          {
            "type": "match",
            "id": "797459ac-1ed1-48db-9c24-f9ddaeefbbc9"
          },
          {
            "type": "match",
            "id": "569b33ab-b3f7-4791-b6a2-7b7bf22cd690"
          },
          {
            "type": "match",
            "id": "1d972810-a85d-4ce1-9837-07c1e6cb68bf"
          },
          {
            "type": "match",
            "id": "79e37122-3b0e-4ed9-acaf-e27b05c3ecb2"
          },
          {
            "type": "match",
            "id": "52ba850d-b062-4ee3-9e36-9cb2f87b3c31"
          },
          {
            "type": "match",
            "id": "5a373ab0-34da-4cbc-9712-1687758229f4"
          },
          {
            "type": "match",
            "id": "7c062590-5ac9-47a5-92c4-299ac154d9af"
          },
          {
            "type": "match",
            "id": "0be96a69-06d2-42bf-adb0-885bb61cc1d9"
          },
          {
            "type": "match",
            "id": "cfeb4e09-65b6-47e0-814d-22dc3c1c6826"
          },
          {
            "type": "match",
            "id": "6f896d3d-55b6-49f2-aeff-3360e7deb398"
          },
          {
            "type": "match",
            "id": "b85f7e71-e9c8-4761-ad31-f3146fd059ce"
          },
          {
            "type": "match",
            "id": "a2dd9289-986d-4e32-8549-166e61f90922"
          },
          {
            "type": "match",
            "id": "30100a35-c09e-464c-99f6-4e11ee6a5ad9"
          },
          {
            "type": "match",
            "id": "655e3f05-bf76-4ad9-8e69-1dc4b6986f7b"
          }
        ]
      }
    }
  }
}
`

func TestParseSamples(t *testing.T) {
	samples, err := ParseSamples(testSamplesStr)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "bluehole-pubg", samples.TitleId)
	assert.Greater(t, len(samples.Matches), 0)
	assert.Equal(t, "b8b0003d-9722-4aa9-ab88-7bdf19c65a3e", samples.Matches[0].ID)
}
