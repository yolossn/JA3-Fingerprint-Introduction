#  JA3-Fingerprint-Introduction

Introduction to JA3 Fingerprint and how to impersonate it using golang.

> Disclaimer: This article is for educational purpose only.

# What is JA3 ? 
  JA3 is a fingerprinting mechanism used to uniquely identify clients based on their TLS clientHello packets. So whenever you access something over the internet your browser/client has to complete a TLS Handshake, this is a multistep process when the client and the server authenticate each other and negotiates security keys to be used for further data transmission.

  JA3 mechanism uses the client Hello packet to create a fingerprint which can be used to identify the operating system and the client from which the request was made. This comes in handy to identify various commonly used malwares and avoid traffic from them to protect your website. 

  Head over to [https://ja3er.com/](https://ja3er.com/), You can find the hash created from the request sent by your browser, and
  click on search for ja3 hash to find your client and operating system. This is mine

  <p align="center">
  <img width="800" height="466" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/ja3er.png">
  </p>


  The site also provides a rest api to find your ja3 fingerprint and infromation about your client and operating system. Lets hit their API via curl see what the output is.

  > curl -X GET 'https://ja3er.com/json'

  This endpoint gives the ja3 and the ja3_hash (fingerprint).
  <p align="center">
  <img width="800" height="120" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/ja3curl1.png">
  </p>


  Get the `ja3_hash` from the previous response and search for it.
  > curl -X GET 'https://ja3er.com/search/3faa4ad39f690c4ef1c3160caa375465'
  
  <p align="center">
  <img width="800" height="160" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/ja3curl2.png">
  </p>


  Don't confuse this with the `User-Agent` which is sent in the request headers. A client can send anything as it is under the control on the client. Even if one sends a different `User-Agent` the source of the request can be found using the JA3
  fingerprint.

  Interesting fact! Have you ever tried to access google using tor browser and been blocked by a screen which says their systems have found unusual traffic, It it done by a mix of fingerprint filtering and by matching IP with list of public tor exit relays. Similar methods are used by govt and other important sites also.

  > Tor Google

  <p align="center">
  <img width="600" height="440" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/torgoogle.png">
  </p>

  > Tor JA3 Fingerprint

  <p align="center">
  <img width="600" height="360" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/torja3.png">
  </p>

  There are also other fingerprinting techniques other than ja3 and it is not public which fingerprinting method is used by Google. They might even have a proprietary technique 🤷‍♂️.

# Impersonating JA3.
  
  So basically if the fingerprint is created using the TLS handshake Client Hello packet then we can create a 
  packet similar to a browser when programatically accessing a website. Yes it is possible with this golang library [JA3Transport](https://github.com/CUCyber/ja3transport). It implements `http.Client` and `http.Transport` interface, allowing us to craft the TLS handshake packets and impersonate someother client. To understand more about the library and ja3 computation logic check the official medium post by the creators of the library. [Link](https://medium.com/cu-cyber/impersonating-ja3-fingerprints-b9f555880e42)

  Let's see the ja3Transport client vs standard golang http client in action.

 ## Normal Http Client
 Using the go http client we hit the same ja3er.com and fetch the results.Check the code [here](https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/NativeHttpClient.go)

  <p align="center">
  <img width="800" height="160" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/NativeClient.png">
  </p>
 


 From the response it is clear that the fingerprint matches highly with `Go-http-client/1.1`.

 ## JA3Transport Http Client
  Now lets use the JA3Transport Http Client and do the same.

  The code changes required is very simple.

  > Create a http client with the ja3 string or predefined browser and use it as your http client.

 ```
  httpClient,err := ja3transport.New(ja3transport.SafariAuto)
	if err != nil{
		fmt.Println(err)
		panic(err)
	}
  ```

  <p align="center">
  <img width="800" height="200" src="https://github.com/yolossn/JA3-Fingerprint-Introduction/blob/master/images/Ja3TransportClient.png">
  </p>

  From the response it is clear the fingerprint has majority of the matches with Mozilla and Safari.

  Hurray 🎉, we successfully impersonated some other client.

  One can also use the [ja3transport.NewWithString](https://godoc.org/github.com/CUCyber/ja3transport#example-NewWithString) method to impersonate a particular client.

  As mentioned earlier the `ja3transport` libarary also implements the `http.Transport` interface which can be used to create proxy server which impersonate some other client. This comes in handy if you have to impersonate traffic from another application or scripts from another programming language.



References:
- [BroCon 2018 - Fingerprinting Encrypted Channels with Bro for High Fidelity Detection](https://www.youtube.com/watch?v=lCa5qyxs-rA)
- [Impersonating JA3 Fingerprints](https://medium.com/cu-cyber/impersonating-ja3-fingerprints-b9f555880e42)
- [Does Google know that I am using Tor Browser?](https://tor.stackexchange.com/questions/313/does-google-know-that-i-am-using-tor-browser)
