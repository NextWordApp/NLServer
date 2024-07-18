# nEXT Word Server

- 用户账户管理
  1. 注册方式
      1. 直接使用邮箱注册，生成对应的ID和密码
      2. 使用未绑定的钱包登录，前后端做验证
      3. 绑定账户信息，github，飞书，抖音，QQ，
  2. 登录方式
      1. 使用 ID/邮箱/github/抖音 密码进行登录
      2. 使用钱包登录
- 背单词交互(用户登录后)
  1. 先检测用户是否登录，若未登录，则使用随机数随机推荐极简词汇
  2. 若登录，
  3. 刷视频逻辑
      1. 则通过用户所学过和未学过的单词来查询用户未学习过的单词
      2. 后续可能是某个算法，比如艾宾浩斯曲线，记忆宫殿规则来查询单词ID
      3. 拿到单词ID后，通过 ID 查询对应单词 信息，给到单词
      4. 再通过 单词 输入 AI(Dify、Coze)，拿到 扩写的相关文本信息
          1. 文生视频
          2. AI 推荐算法推荐已有的感兴趣的视频
          3. ChatTTS 中文转英文
      5. 将 中文，英文，单词信息，视频 URL 均返回给前端
      6. 前端记录用户滑动时间，退出时间，进入时间，是否喜欢，并发送给后端
          1. 后端可以根据这些内容来个性化推荐用户喜欢的视频给到对应用户
            
