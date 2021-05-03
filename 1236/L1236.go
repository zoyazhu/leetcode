package main
/*
1236. 网络爬虫
给定一个链接 startUrl 和一个接口 HtmlParser ，请你实现一个网络爬虫，以实现爬取同 startUrl 拥有相同 域名标签 的全部链接。
该爬虫得到的全部链接可以 任何顺序 返回结果。

你的网络爬虫应当按照如下模式工作：
自链接 startUrl 开始爬取
调用 HtmlParser.getUrls(url) 来获得链接url页面中的全部链接
同一个链接最多只爬取一次
只输出 域名 与 startUrl 相同 的链接集合

如上所示的一个链接，其域名为 example.org。简单起见，你可以假设所有的链接都采用 http协议 并没有指定 端口。
例如，链接 http://leetcode.com/problems 和 http://leetcode.com/contest 是同一个域名下的，
而链接http://example.org/test 和 http://example.com/abc 是不在同一域名下的。

HtmlParser 接口定义如下：
interface HtmlParser {
  // 返回给定 url 对应的页面中的全部 url 。
  public List<String> getUrls(String url);
}
下面是两个实例，用以解释该问题的设计功能，对于自定义测试，你可以使用三个变量  urls, edges 和 startUrl。注意在代码实现中，
你只可以访问 startUrl ，而 urls 和 edges 不可以在你的代码中被直接访问。

示例 1：
输入：
urls = [
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.google.com",
  "http://news.yahoo.com/us"
]
edges = [[2,0],[2,1],[3,2],[3,1],[0,4]]
startUrl = "http://news.yahoo.com/news/topics/"
输出：[
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.yahoo.com/us"
]

示例 2：
输入：
urls = [
  "http://news.yahoo.com",
  "http://news.yahoo.com/news",
  "http://news.yahoo.com/news/topics/",
  "http://news.google.com"
]
edges = [[0,2],[2,1],[3,2],[3,1],[3,0]]
startUrl = "http://news.google.com"
输入：["http://news.google.com"]
解释：startUrl 链接到所有其他不共享相同主机名的页面。
*/

/**
 * // This is HtmlParser's API interface.
 * // You should not implement it, or speculate about its implementation
 * type HtmlParser struct {
 *     func GetUrls(url string) []string {}
 * }
 */
type HtmlParser struct {
	//func GetUrls(url string) []string {}
}


func crawl(startUrl string, htmlParser HtmlParser) []string {
	return nil
}

/*
解题思路
其实本题就是一个简单的层序遍历问题，用 BFS 即可解决，但是题目描述比较晦涩。

题目解读：

1、给定一个 startUrl，这个 startUrl 的页面存在若干个链接，可以用 HtmlParser 接口的 getUrls(String url) 方法获取该页面的所有链接；

2、返回能够访问到的、和 startUrl 具有同样的主机名的所有链接。

本题的 BFS 逻辑中涉及到“查看当前链接是否已使用”的操作，建议使用 HashSet 来临时存储，到需要返回的时候再把 set 的数据装入 list 中，可以提升 10 倍的速度，原因在于 HashSet 的查找速度比 ArrayList 快。

代


import java.util.*;

class Solution {

public List<String> crawl(String startUrl, HtmlParser htmlParser) {
Queue<String> queue = new LinkedList<>();
// 这里用 HashSet 比直接用 ArrayList 快十倍，到最后再装入 list
HashSet<String> set = new HashSet<>();
set.add(startUrl);
queue.add(startUrl);
while (!queue.isEmpty()) {
String currentUrl = queue.poll();
List<String> nextUrls = htmlParser.getUrls(currentUrl);
for (String nextUrl : nextUrls) {
if (!set.contains(nextUrl) && getHostName(nextUrl).equals(getHostName(currentUrl))) {
// 如果链接没使用过且和 currentUrl 同主机，则有效
set.add(nextUrl);
queue.add(nextUrl);
}
}
}
return new ArrayList<>(set);
}

private String getHostName(String url) {
url = url.substring(7);
if (url.contains("/")) {
int end = url.indexOf('/'); // 第一次出现斜杠的索引位置
return url.substring(0, end);
} else {
return url;
}
}
}

作者：klb
链接：https://leetcode-cn.com/problems/web-crawler/solution/1236-wang-luo-pa-chong-by-klb/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

 */
