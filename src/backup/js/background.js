
const tasks = [
  {id: 1, name: 'One', done: true},
  {id: 2, name: 'Two Two', done: false},
  {id: 3, name: 'Three Three Three', done: true},
  {id: 4, name: 'four four four four', done: true},
]

chrome.runtime.onInstalled.addListener(function() {
  chrome.storage.sync.set({tasks: tasks }, function() {
    console.log("Task setting")
  })
  chrome.declarativeContent.onPageChanged.removeRules(undefined, function() {      
    chrome.declarativeContent.onPageChanged.addRules([{
      conditions: [new chrome.declarativeContent.PageStateMatcher({
        pageUrl: {hostEquals: 'developer.chrome.com'},
      })
      ],
          actions: [new chrome.declarativeContent.ShowPageAction()]
    }]);      
  });
})  
  
  /*
  chrome.runtime.onInstalled.addListener(function() {
    chrome.storage.sync.set({color: '#3aa757'}, function() {
      console.log("The color is green.");
    });
    chrome.declarativeContent.onPageChanged.removeRules(undefined, function() {      
      chrome.declarativeContent.onPageChanged.addRules([{
        conditions: [new chrome.declarativeContent.PageStateMatcher({
          pageUrl: {hostEquals: 'developer.chrome.com'},
        })
        ],
            actions: [new chrome.declarativeContent.ShowPageAction()]
      }]);      
    });
  });
*/