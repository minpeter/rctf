import ctftime from "./util/ctftime";

let config;

try {
  config = JSON.parse(
    document.head.querySelector("meta[name=rctf-config]").content
  );
} catch (error) {
  console.error("Error parsing JSON from meta tag:", error);
  console.log("Using default config");
  config = {
    ctftime: {
      clientId: "0f18c7f2d534b23a56c8",
    },
    meta: {
      description: "A description of your CTF",
      imageUrl: "https: //example.com",
    },
    homeContent: "A description of your CTF. Markdown supported.",
    sponsors: [],
    globalSiteTag: "undefined",
    ctfName: "rCTF",
    divisions: {
      open: "Open",
    },
    defaultDivision: "undefined",
    origin: "http: //localhost:3000",
    startTime: 1702355705000,
    endTime: 0,
    emailEnabled: false,
    userMembers: true,
    faviconUrl: "https://zany.sh/favicon.svg?emoji=🤑",
  };
}

export default config;