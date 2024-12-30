using dotenv.net;

using Scriban;

using Willpinha.Console;

// Don't forget to create a .env file based on the .env.example file during development
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

var templateContent = await File.ReadAllTextAsync("Willpinha.Console/Templates/README.scriban");
var template = Template.Parse(templateContent);

var result = await template.RenderAsync(new
{
    lastUpdated = DateTime.UtcNow.ToString("MM/dd/yyyy HH:mm:ss zzz"),
    famousRepositories = famousRepositories.Items,
    latestPullRequests = latestPullRequests.Items,
    latestIssues = latestIssues.Items
});

var target = args.Any(x => x == "OVERRIDE_README") ? "README.md" : "README.test.md";

await File.WriteAllTextAsync(target, result);

