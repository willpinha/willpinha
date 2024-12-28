using dotenv.net;

using Willpinha.Console;

DotEnv.Load();

// Necessary to avoid rate limit
await Client.Instance.User.Get(Requests.Username);

var apiInfo = Client.Instance.GetLastApiInfo();

if (apiInfo.RateLimit.Remaining == 0)
{
    throw new Exception("GitHub API rate limit reached");
}

Console.WriteLine($"Rate limit: {apiInfo.RateLimit.Remaining}/{apiInfo.RateLimit.Limit}");

var famousRepositories = await Requests.GetFamousRepositories();
var latestPullRequests = await Requests.GetLatestPullRequests();
var latestIssues = await Requests.GetLatestIssues();

Console.WriteLine("Famous repositories:");
foreach (var repo in famousRepositories.Items)
{
    Console.WriteLine($"- {repo.Name} ({repo.StargazersCount} stars)");
}

Console.WriteLine("Latest pull requests:");
foreach (var issue in latestPullRequests.Items)
{
    Console.WriteLine($"- {issue.Title} ({issue.UpdatedAt}) {issue.HtmlUrl}");
}

Console.WriteLine("Latest issues:");
foreach (var issue in latestIssues.Items)
{
    Console.WriteLine($"- {issue.Title} ({issue.UpdatedAt}) {issue.HtmlUrl}");
}