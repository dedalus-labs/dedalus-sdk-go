# Changelog

## 0.1.0-alpha.4 (2025-12-17)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/dedalus-labs/dedalus-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* **api:** config update for dedalus-ai/dev ([af41fea](https://github.com/dedalus-labs/dedalus-sdk-go/commit/af41fea3fa08b1c8c0595ae7346c0c525be5ae24))
* **api:** config update for dedalus-ai/dev ([d240153](https://github.com/dedalus-labs/dedalus-sdk-go/commit/d240153d2233d17c0cd9a1407e4a15b8e2816528))
* **api:** improve types ([981e302](https://github.com/dedalus-labs/dedalus-sdk-go/commit/981e302628d2adcf5448289b7834746f142c1b26))
* **api:** mcp server params ([3417509](https://github.com/dedalus-labs/dedalus-sdk-go/commit/34175093b827bf28d333fccc6b46d233a5a375d0))
* **api:** messages param nullable ([4408331](https://github.com/dedalus-labs/dedalus-sdk-go/commit/4408331c403387ad0b0e944dcc4141e1d5c3eafa))
* **api:** response format ([87661ed](https://github.com/dedalus-labs/dedalus-sdk-go/commit/87661ede914cd98aeeacb141a533292623beedc6))
* **api:** schema compiler landed ([77a9162](https://github.com/dedalus-labs/dedalus-sdk-go/commit/77a9162eacfa2f4aedfd3cc9a144c0a85fcac482))
* **api:** standardize name casing with stainless initialism ([4291ed2](https://github.com/dedalus-labs/dedalus-sdk-go/commit/4291ed2e093ed660295f87992eec81bf5479012f))


### Bug Fixes

* **api:** add thought signature ([3320501](https://github.com/dedalus-labs/dedalus-sdk-go/commit/3320501ff18f520c131398a789e424f6540d185e))
* **api:** mcp credential types ([294857c](https://github.com/dedalus-labs/dedalus-sdk-go/commit/294857cdd8ece556942094e93f4971a52b914fa9))
* **mcp:** correct code tool API endpoint ([10eac06](https://github.com/dedalus-labs/dedalus-sdk-go/commit/10eac06866c6bc23286992a33ed20cd0f5b27cb7))
* rename param to avoid collision ([9dfa024](https://github.com/dedalus-labs/dedalus-sdk-go/commit/9dfa024195bd09e018fc929d28872d3605f9a0a3))
* skip usage tests that don't work with Prism ([2aab71d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/2aab71d2cc3ed8e7f67bbe9d3400295b0e44514d))


### Chores

* **api:** migrate pkg manager to uv ([a3152f2](https://github.com/dedalus-labs/dedalus-sdk-go/commit/a3152f2f36b47e8ddf3a219e67a8516fe34ee384))
* **api:** point local dev to 4010 port for prism ([7139c3e](https://github.com/dedalus-labs/dedalus-sdk-go/commit/7139c3e82d68d21ad6bd5ddfcf764caa939b87cf))
* **auth:** add minor auth params ([bd1f962](https://github.com/dedalus-labs/dedalus-sdk-go/commit/bd1f96277a36b71c496d3722a708fc0fa02f1e8d))
* bump gjson version ([aef2b2d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/aef2b2de5b5d9c540c15313b88dc98a6118a2ed0))
* elide duplicate aliases ([45e32df](https://github.com/dedalus-labs/dedalus-sdk-go/commit/45e32df175a4c4c25cc6a443bc7ca7de1946a99f))
* **internal:** codegen related update ([99a5aae](https://github.com/dedalus-labs/dedalus-sdk-go/commit/99a5aae0ad7f6fd85d5f865094570038ca9eb9dd))


### Refactors

* **api:** types for mcp server serialization ([bee352c](https://github.com/dedalus-labs/dedalus-sdk-go/commit/bee352ca340111af38b261b930631c7e24c7e0c1))
* **api:** update auth types ([816dff3](https://github.com/dedalus-labs/dedalus-sdk-go/commit/816dff340e9b0fb0481f85da4d26e1ebd439490b))

## 0.1.0-alpha.3 (2025-11-08)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/dedalus-labs/dedalus-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** add endpoints ([56c793d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/56c793dc4e5cca9c67ad26938837d3b83c76c8ae))
* **api:** add endpoints ([2e8fbb2](https://github.com/dedalus-labs/dedalus-sdk-go/commit/2e8fbb25bbf9bf0a077c4c958ebe0fc564c15321))
* **api:** adjust parameters ([0219a27](https://github.com/dedalus-labs/dedalus-sdk-go/commit/0219a279ae81f19a6ff44d39d7842f57179ec711))
* **api:** api update ([6c33e7f](https://github.com/dedalus-labs/dedalus-sdk-go/commit/6c33e7f4be92289ddf9a3d90aa56e0a47e0d46dd))
* **api:** api update ([ca846dc](https://github.com/dedalus-labs/dedalus-sdk-go/commit/ca846dc4613bd575a30ecd34a404eba62f2b07c7))
* **api:** api update ([42eb47f](https://github.com/dedalus-labs/dedalus-sdk-go/commit/42eb47f63f91eb91fbd764c423f54a904394a350))
* **api:** api update ([74a33ec](https://github.com/dedalus-labs/dedalus-sdk-go/commit/74a33ec6490bc81431ec8b3eef7573f5ccb25ae5))
* **api:** api update ([473af2c](https://github.com/dedalus-labs/dedalus-sdk-go/commit/473af2ce0724e9c5ae85fcaec6fe5222a1bfd316))
* **api:** auto exec tools ([ad5e300](https://github.com/dedalus-labs/dedalus-sdk-go/commit/ad5e30016e7729d1d611b788d36ba183409286a8))
* **api:** chat completions ([c0b25f6](https://github.com/dedalus-labs/dedalus-sdk-go/commit/c0b25f646add6b4b968c4989f03d3c7f7ca77dc7))
* **api:** Config update for dedalus-ai/dev ([5046221](https://github.com/dedalus-labs/dedalus-sdk-go/commit/50462218937491963dce3ce8c3218e594f842a2b))
* **api:** decouple Model and DedalusModel ([a42e224](https://github.com/dedalus-labs/dedalus-sdk-go/commit/a42e224f42e4423c1b453ea3626403e61f196cc1))
* **api:** dedalus model update ([8064352](https://github.com/dedalus-labs/dedalus-sdk-go/commit/80643523a4ecb9928654c867fa14db5789a95cec))
* **api:** fixing streaming again ([5501f65](https://github.com/dedalus-labs/dedalus-sdk-go/commit/5501f65d2a32b97c4e7d3c2e68dbb23dc11291bd))
* **api:** id-&gt;name in DedalusModel ([1f8597d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/1f8597dffac0cd91383fecdc9f2c3d2d255218dc))
* **api:** image support ([56093bf](https://github.com/dedalus-labs/dedalus-sdk-go/commit/56093bfd630bfefc3ec24f1d324e1f8ac028eaf0))
* **api:** logic adj ([c432afe](https://github.com/dedalus-labs/dedalus-sdk-go/commit/c432afe6ff1e05dd9429b8c6937717f04b8552b6))
* **api:** manual updates ([8a3df23](https://github.com/dedalus-labs/dedalus-sdk-go/commit/8a3df230e05b07d768731667cc58f46e954d6239))
* **api:** manual updates ([b4acd12](https://github.com/dedalus-labs/dedalus-sdk-go/commit/b4acd129b6d48ee5018d970fe12eb194dee7a242))
* **api:** ModelConfig ([26935d4](https://github.com/dedalus-labs/dedalus-sdk-go/commit/26935d4a6cfd5567b6411682e4867d55973f9f7f))
* **api:** polished types ([7886b10](https://github.com/dedalus-labs/dedalus-sdk-go/commit/7886b105133b4e6f2342412e8f9053903a390d13))
* **api:** spec concise ([7282fb1](https://github.com/dedalus-labs/dedalus-sdk-go/commit/7282fb16cfdf441da6ade6373bb6a9aa24aed64a))
* **api:** streaming change ([8382b9f](https://github.com/dedalus-labs/dedalus-sdk-go/commit/8382b9f25926c3b87816bbb1ec2fb74301023576))
* **api:** streaming schemas ([349c63e](https://github.com/dedalus-labs/dedalus-sdk-go/commit/349c63eeea3b5ea33264a269d8c96515bca1c6e8))
* **api:** to_schema and Model class ([2322e45](https://github.com/dedalus-labs/dedalus-sdk-go/commit/2322e457d492eb5718586885dee3c12d710c3ae6))
* **api:** update types ([c935f94](https://github.com/dedalus-labs/dedalus-sdk-go/commit/c935f946641b70d70a6c4ee15afd0741e39cea37))
* **client:** support optional json html escaping ([a7fb2b4](https://github.com/dedalus-labs/dedalus-sdk-go/commit/a7fb2b4346e772d737562dfd713cb041385ee3f1))
* **model:** add DedalusModel ([ba9ca74](https://github.com/dedalus-labs/dedalus-sdk-go/commit/ba9ca74e89e709d9f62a9708a91b89dd694cded8))


### Bug Fixes

* **api:** add shared DedalusModel type ([9ab91d3](https://github.com/dedalus-labs/dedalus-sdk-go/commit/9ab91d31ad5563224d96beb33c082ad8f4e42a4a))
* bugfix for setting JSON keys with special characters ([84abc7a](https://github.com/dedalus-labs/dedalus-sdk-go/commit/84abc7a0782b7d20912578b0dcda5ececd2cac4a))
* close body before retrying ([c9078f2](https://github.com/dedalus-labs/dedalus-sdk-go/commit/c9078f27419b32af430c6f902a49a37cbe82821b))
* **internal:** unmarshal correctly when there are multiple discriminators ([50bb19d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/50bb19d8bafb5d1c96b892dc6f2d2844915f77c5))
* remove null from release please manifest ([03a3db0](https://github.com/dedalus-labs/dedalus-sdk-go/commit/03a3db077be20f3bf01a26dcfa983221d463f9c3))
* use release please annotations on more places ([7bfe9e8](https://github.com/dedalus-labs/dedalus-sdk-go/commit/7bfe9e89bea66513f00aecd28573b84db8f122ca))
* use slices.Concat instead of sometimes modifying r.Options ([c23ec78](https://github.com/dedalus-labs/dedalus-sdk-go/commit/c23ec78490fe777b4a4fb35a2f7b079560cbc014))


### Chores

* bump minimum go version to 1.22 ([7cba6f1](https://github.com/dedalus-labs/dedalus-sdk-go/commit/7cba6f1d2f13ce7d20f97bb6032fd910decd60c5))
* do not install brew dependencies in ./scripts/bootstrap by default ([d49d734](https://github.com/dedalus-labs/dedalus-sdk-go/commit/d49d7349d4e432fe947410697e3a2b9d89a41fbb))
* **internal:** codegen related update ([4b5c811](https://github.com/dedalus-labs/dedalus-sdk-go/commit/4b5c8119c41b2277d58cd629515557a0f40797e2))
* **internal:** codegen related update ([ec26673](https://github.com/dedalus-labs/dedalus-sdk-go/commit/ec26673e94b1da36d7920805c1ea202088d563a8))
* **internal:** grammar fix (it's -&gt; its) ([56784ab](https://github.com/dedalus-labs/dedalus-sdk-go/commit/56784ab9f63dcee7a6c71d45daa8b40c5bb95933))
* **internal:** update comment in script ([a684b4d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/a684b4ddcf252a09c841083ffd6a538bfa615e2c))
* update @stainless-api/prism-cli to v5.15.0 ([2a8068d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/2a8068daa10912b171f94c737d7898ae2d91374e))
* update more docs for 1.22 ([3623d51](https://github.com/dedalus-labs/dedalus-sdk-go/commit/3623d513c53add079893ce773a4c8580f9865938))

## 0.1.0-alpha.2 (2025-08-05)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/dedalus-labs/dedalus-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Features

* **api:** add streaming ([67a460d](https://github.com/dedalus-labs/dedalus-sdk-go/commit/67a460d692bba17c1a86417ac932206ca556d38a))
* **api:** add streaming configuration ([6136a95](https://github.com/dedalus-labs/dedalus-sdk-go/commit/6136a95170f0cc3927c205024f955cc8c5d5f267))
* **api:** api update ([497edff](https://github.com/dedalus-labs/dedalus-sdk-go/commit/497edff59925bb26099cd715325aa9a1ae7a95a4))
* **api:** revert streaming for now ([1b58e24](https://github.com/dedalus-labs/dedalus-sdk-go/commit/1b58e24daf9dbde766d164e619916207087bc471))

## 0.1.0-alpha.1 (2025-07-30)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/dedalus-labs/dedalus-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update via SDK Studio ([5ffa589](https://github.com/dedalus-labs/dedalus-sdk-go/commit/5ffa589805c4e69331917b4bf641a7a598c9b7e3))


### Chores

* configure new SDK language ([79e69c8](https://github.com/dedalus-labs/dedalus-sdk-go/commit/79e69c8e6b6de1ef738efe6824656357b1cebdae))
* update SDK settings ([e15f7f0](https://github.com/dedalus-labs/dedalus-sdk-go/commit/e15f7f0d60e284983390e7fd3ed56d1aefd581cb))
