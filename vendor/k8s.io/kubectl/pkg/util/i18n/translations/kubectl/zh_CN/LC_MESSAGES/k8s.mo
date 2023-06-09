Þ    f      L     |         z   ¡  ´   	  <   Ñ	  S   
  <   b
  c  
  ´    .   ¸  "   ç  4   
     ?     \    {  X     o   Ú    J  v   L  t   Ã  Ä  8  ;   ý  [   9  J     a   à  ½   B  Å      ®   Æ  %   u  W        ó  u     4     -   ¼  3   ê  2        Q  *   e  .     *   ¿  0   ê  0     0   L  "   }        *   ¾  A   é     +  )   I     s           «  (   Ì     õ  `        m     	     ¡     º  $   Õ     ú       a   0  s     B     +   I  +   u  6   ¡  q   Ø  /   J   1   z   '   ¬      Ô   &   í   %   !  (   :!  #   c!      !     ¨!  9   Ç!     "      "  #   :"     ^"  H   á"  &   *#  e   Q#  å   ·#  E   $  a   ã$  Ë   E%  Ï   &     á&     þ&  =   '  $   T'     y'  &   '  +   À'     ì'  r   (     t(  /   (    ¸(  x   ;*     ´*  8   @+  A   y+  6   »+  a  ò+     T-  -   õ.  -   #/  0   Q/     /  !   ¡/  í   Ã/  a   ±0  r   1  ×   1     ^2  {   å2  Û  a3  C   =5  Z   5  F   Ü5  Z   #6  ¨   ~6  Á   '7  ¢   é7     8  \   «8  !   9  m   *9  +   9  $   Ä9  /   é9  B   :     \:  *   t:  0   :  +   Ð:  '   ü:  '   $;  *   L;     w;     ;  '   µ;  E   Ý;     #<     ?<     [<  &   w<  $   <  1   Ã<     õ<  ^   =     p=     ù=     ~>     >  '   ³>     Û>     ñ>  o   
?  v   z?  =   ñ?  ,   /@  %   \@  *   @  `   ­@  *   A  !   9A     [A     yA      A     °A  4   ÏA  "   B     'B     CB  -   \B     B      B     »B  f   ÛB  ;   BC     ~C  Q   C  è   ïC  :   ØD  N   E  °   bE  °   F     ÄF     ßF  3   øF     ,G     HG  *   dG  /   G     ¿G  u   ÒG     HH  )   \H     	   H       -   #                  3      `       d                  C          I       A          1           >   0          !           "   (       L   %       5   J   ?   4   )   b   Z   @      f   F       =         ;              c   ^         9   [   e   M      a   ,      S   '      \          Q             .   V   T   W       B      Y          E      6      :      X   &   /       P       D   K      U   2                      _   7   ]   <   8           R          $      G              O   N   *   
   +    
		  # Show metrics for all nodes
		  kubectl top node

		  # Show metrics for a given node
		  kubectl top node NODE_NAME 
		# Get the documentation of the resource and its fields
		kubectl explain pods

		# Get the documentation of a specific field of a resource
		kubectl explain pods.spec.containers 
		# Print flags inherited by all commands
		kubectl options 
		# Print the client and server versions for the current context
		kubectl version 
		# Print the supported API versions
		kubectl api-versions 
		# Show metrics for all pods in the default namespace
		kubectl top pod

		# Show metrics for all pods in the given namespace
		kubectl top pod --namespace=NAMESPACE

		# Show metrics for a given pod and its containers
		kubectl top pod POD_NAME --containers

		# Show metrics for the pods defined by label name=myLabel
		kubectl top pod -l name=myLabel 
		Convert config files between different API versions. Both YAML
		and JSON formats are accepted.

		The command takes filename, directory, or URL as input, and convert it into format
		of version specified by --output-version flag. If target version is not specified or
		not supported, convert to latest version.

		The default output will be printed to stdout in YAML format. One can use -o option
		to change to output destination. 
		Create a namespace with the specified name. 
		Create a role with single rule. 
		Create a service account with the specified name. 
		Mark node as schedulable. 
		Mark node as unschedulable. 
		Set the latest last-applied-configuration annotations by setting it to match the contents of a file.
		This results in the last-applied-configuration being updated as though 'kubectl apply -f <file>' was run,
		without updating any other parts of the object. 
	  # Create a new namespace named my-namespace
	  kubectl create namespace my-namespace 
	  # Create a new service account named my-service-account
	  kubectl create serviceaccount my-service-account 
	Create an ExternalName service with the specified name.

	ExternalName service references to an external DNS address instead of
	only pods, which will allow application authors to reference services
	that exist off platform, on other clusters, or locally. 
	Help provides help for any command in the application.
	Simply type kubectl help [path to command] for full details. 
    # Create a new LoadBalancer service named my-lbs
    kubectl create service loadbalancer my-lbs --tcp=5678:8080 
    # Dump current cluster state to stdout
    kubectl cluster-info dump

    # Dump current cluster state to /path/to/cluster-state
    kubectl cluster-info dump --output-directory=/path/to/cluster-state

    # Dump all namespaces to stdout
    kubectl cluster-info dump --all-namespaces

    # Dump a set of namespaces to /path/to/cluster-state
    kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state 
    Create a LoadBalancer service with the specified name. A comma-delimited set of quota scopes that must all match each object tracked by the quota. A comma-delimited set of resource=quantity pairs that define a hard limit. A label selector to use for this budget. Only equality-based selector requirements are supported. A label selector to use for this service. Only equality-based selector requirements are supported. If empty (the default) infer the selector from the replication controller or replica set.) Additional external IP address (not managed by Kubernetes) to accept for the service. If this IP is routed to a node, the service can be accessed by this IP in addition to its generated service IP. An inline JSON override for the generated object. If this is non-empty, it is used to override the generated object. Requires that the object supply a valid apiVersion field. Approve a certificate signing request Assign your own ClusterIP or set to 'None' for a 'headless' service (no loadbalancing). Attach to a running container ClusterIP to be assigned to the service. Leave empty to auto-allocate, or set to 'None' to create a headless service. ClusterRole this ClusterRoleBinding should reference ClusterRole this RoleBinding should reference Convert config files between different API versions Copy files and directories to and from containers. Create a TLS secret Create a namespace with the specified name Create a secret for use with a Docker registry Create a secret using specified subcommand Create a service account with the specified name Delete the specified cluster from the kubeconfig Delete the specified context from the kubeconfig Deny a certificate signing request Describe one or many contexts Display clusters defined in the kubeconfig Display merged kubeconfig settings or a specified kubeconfig file Display one or many resources Drain node in preparation for maintenance Edit a resource on the server Email for Docker registry Execute a command in a container Forward one or more local ports to a pod Help about any command If non-empty, set the session affinity for the service to this; legal values: 'None', 'ClientIP' If non-empty, the annotation update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. If non-empty, the labels update will only succeed if this is the current resource-version for the object. Only valid when specifying a single resource. Mark node as schedulable Mark node as unschedulable Mark the provided resource as paused Modify certificate resources. Modify kubeconfig files Name or number for the port on the container that the service should direct traffic to. Optional. Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time / since may be used. Output shell completion code for the specified shell (bash or zsh) Password for Docker registry authentication Path to PEM encoded public key certificate. Path to private key associated with given certificate. Precondition for resource version. Requires that the current resource version match this value in order to scale. Print the client and server version information Print the list of flags inherited by all commands Print the logs for a container in a pod Resume a paused resource Role this RoleBinding should reference Run a particular image on the cluster Run a proxy to the Kubernetes API server Server location for Docker registry Set specific features on objects Set the selector on a resource Show details of a specific resource or group of resources Show the status of the rollout Synonym for --target-port The image for the container to run. The image pull policy for the container. If left empty, this value will not be specified by the client and defaulted by the server The minimum number or percentage of available pods this budget requires. The name for the newly created object. The name for the newly created object. If not specified, the name of the input resource will be used. The name of the API generator to use. There are 2 generators: 'service/v1' and 'service/v2'. The only difference between them is that service port in v1 is named 'default', while it is left unnamed in v2. Default is 'service/v2'. The network protocol for the service to be created. Default is 'TCP'. The port that the service should serve on. Copied from the resource being exposed, if unspecified The resource requirement limits for this container.  For example, 'cpu=200m,memory=512Mi'.  Note that server side components may assign limits depending on the server configuration, such as limit ranges. The resource requirement requests for this container.  For example, 'cpu=100m,memory=256Mi'.  Note that server side components may assign requests depending on the server configuration, such as limit ranges. The type of secret to create Undo a previous rollout Update resource requests/limits on objects with pod templates Update the annotations on a resource Update the labels on a resource Update the taints on one or more nodes Username for Docker registry authentication View rollout history Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory dummy restart flag) kubectl controls the Kubernetes cluster manager Project-Id-Version: gettext-go-examples-hello
Report-Msgid-Bugs-To: EMAIL
PO-Revision-Date: 2022-07-04 18:54+0800
Last-Translator: zhengjiajin <zhengjiajin@caicloud.io>
Language-Team: 
Language: zh
MIME-Version: 1.0
Content-Type: text/plain; charset=UTF-8
Content-Transfer-Encoding: 8bit
Plural-Forms: nplurals=2; plural=(n > 1);
X-Generator: Poedit 3.0.1
X-Poedit-SourceCharset: UTF-8
 
		  # æ¾ç¤ºææèç¹çææ 
		  kubectl top ode

		  # æ¾ç¤ºæå®èç¹çææ 
		  kubectl top node NODE_NAME 
		# è·åèµæºåå¶å­æ®µçææ¡£
		kubectl explain pods

		# è·åèµæºæå®å­æ®µçææ¡£
		kubectl explain pods.spec.containers 
		# è¾åºææå½ä»¤ç»§æ¿ç flags
		kubectl options 
		# è¾åºå½åå®¢æ·ç«¯åæå¡ç«¯ççæ¬
		kubectl version 
		# è¾åºæ¯æç API çæ¬
		kubectl api-versions 
		# æ¾ç¤º default å½åç©ºé´ä¸ææ Pods çææ 
		kubectl top pod

		# æ¾ç¤ºæå®å½åç©ºé´ä¸ææ Pods çææ 
		kubectl top pod --namespace=NAMESPACE

		# æ¾ç¤ºæå® Pod åå®çå®¹å¨ç metrics
		kubectl top pod POD_NAME --containers

		# æ¾ç¤ºæå® label ä¸º name=myLabel ç Pods ç metrics
		kubectl top pod -l name=myLabel 
		å¨ä¸åç API çæ¬ä¹é´è½¬æ¢éç½®æä»¶ãæ¥å YAML
		å JSON æ ¼å¼ã

		è¿ä¸ªå½ä»¤ä»¥æä»¶å, ç®å½, æè URL ä½ä¸ºè¾å¥ï¼å¹¶éè¿ âoutput-version åæ°
		 è½¬æ¢å°æå®çæ¬çæ ¼å¼ãå¦ææ²¡ææå®ç®æ çæ¬æèææå®çæ¬
		ä¸æ¯æ, åè½¬æ¢ä¸ºææ°çæ¬ã

		é»è®¤ä»¥ YAML æ ¼å¼è¾åºå°æ åè¾åºãå¯ä»¥ä½¿ç¨ -o option
		ä¿®æ¹ç®æ è¾åºçæ ¼å¼ã 
		ç¨ç»å®åç§°åå»ºä¸ä¸ªå½åç©ºé´ã 
		åå»ºä¸ä¸ªå·æåä¸è§åçè§è²ã 
		ç¨æå®çåç§°åå»ºä¸ä¸ªæå¡è´¦æ·ã 
		æ è®°èç¹ä¸ºå¯è°åº¦ã 
		æ è®°èç¹ä¸ºä¸å¯è°åº¦ã 
		è®¾ç½®ææ°ç last-applied-configuration æ³¨è§£ï¼ä½¿ä¹å¹éææä»¶çåå®¹ã
		è¿ä¼å¯¼è´ last-applied-configuration è¢«æ´æ°ï¼å°±åæ§è¡äº kubectl apply -f <file> ä¸æ ·ï¼
		åªæ¯ä¸ä¼æ´æ°å¯¹è±¡çå¶ä»é¨åã 
	  # åå»ºä¸ä¸ªåä¸º my-namespace çæ°å½åç©ºé´
	  kubectl create namespace my-namespace 
	  # åå»ºä¸ä¸ªåä¸º my-service-account çæ°æå¡å¸æ·
	  kubectl create serviceaccount my-service-account 
	åå»ºå·ææå®åç§°ç ExternalName æå¡ã

	ExternalName æå¡å¼ç¨å¤é¨ DNS å°åèä¸æ¯ Pod å°åï¼
	è¿å°åè®¸åºç¨ç¨åºä½èå¼ç¨å­å¨äºå¹³å°å¤ãå¶ä»éç¾¤ä¸ææ¬å°çæå¡ã 
	Help ä¸ºåºç¨ç¨åºä¸­çä»»ä½å½ä»¤æä¾å¸®å©ã
	åªéé®å¥ kubectl help [å½ä»¤è·¯å¾] å³å¯è·å¾å®æ´çè¯¦ç»ä¿¡æ¯ã 
    # åå»ºä¸ä¸ªåç§°ä¸º my-lbs çæ°è´è½½åè¡¡æå¡
    kubectl create service loadbalancer my-lbs --tcp=5678:8080 
    # å¯¼åºå½åçéç¾¤ç¶æä¿¡æ¯å°æ åè¾åº
    kubectl cluster-info dump

    # å¯¼åºå½åçéç¾¤ç¶æå° /path/to/cluster-state
    kubectl cluster-info dump --output-directory=/path/to/cluster-state

    # å¯¼åºææå½åç©ºé´å°æ åè¾åº
    kubectl cluster-info dump --all-namespaces

    # å¯¼åºä¸ç»å½åç©ºé´å° /path/to/cluster-state
    kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state 
    ä½¿ç¨ä¸ä¸ªæå®çåç§°åå»ºä¸ä¸ª LoadBalancer æå¡ã ä¸ç»ä»¥éå·åéçéé¢èå´ï¼å¿é¡»å¨é¨å¹ééé¢æè·è¸ªçæ¯ä¸ªå¯¹è±¡ã ä¸ç»ä»¥éå·åéçèµæº=æ°éå¯¹ï¼ç¨äºå®ä¹ç¡¬æ§éå¶ã ä¸ä¸ªç¨äºè¯¥é¢ç®çæ ç­¾éæ©å¨ãåªæ¯æåºäºç­å¼æ¯è¾çéæ©å¨è¦æ±ã ç¨äºæ­¤æå¡çæ ç­¾éæ©å¨ãä»æ¯æåºäºç­å¼æ¯è¾çéæ©å¨è¦æ±ãå¦æä¸ºç©ºï¼é»è®¤ï¼ï¼åä»å¯æ¬æ§å¶å¨æå¯æ¬éä¸­æ¨æ­éæ©å¨ãï¼ ä¸ºæå¡ææ¥åçå¶ä»å¤é¨ IP å°åï¼ä¸ç± Kubernetes ç®¡çï¼ãå¦æè¿ä¸ª IP è¢«è·¯ç±å°ä¸ä¸ªèç¹ï¼é¤äºå¶çæçæå¡ IP å¤ï¼è¿å¯ä»¥éè¿è¿ä¸ª IP è®¿é®æå¡ã éå¯¹æçæå¯¹è±¡çåè JSON è¦çãå¦æè¿ä¸å¯¹è±¡æ¯éç©ºçï¼å°ç¨äºè¦çæçæçå¯¹è±¡ãè¦æ±å¯¹è±¡æä¾ææç apiVersion å­æ®µã æ¹åä¸ä¸ªè¯ä¹¦ç­¾ç½²è¯·æ± ä¸ºâæ å¤´âæå¡ï¼æ è´è½½å¹³è¡¡ï¼åéä½ èªå·±ç ClusterIP æè®¾ç½®ä¸ºâæ ã ææ¥å°ä¸ä¸ªè¿è¡ä¸­çå®¹å¨ è¦åéç»æå¡ç ClusterIPãçç©ºè¡¨ç¤ºèªå¨åéï¼æè®¾ç½®ä¸º âNoneâ ä»¥åå»ºæ å¤´æå¡ã ClusterRoleBinding åºè¯¥æå® ClusterRole RoleBinding åºè¯¥æå® ClusterRole å¨ä¸åç API çæ¬ä¹é´è½¬æ¢éç½®æä»¶ å°æä»¶åç®å½å¤å¶å°å®¹å¨ä¸­æä»å®¹å¨ä¸­å¤å¶åºæ¥ã åå»ºä¸ä¸ª TLS secret ç¨æå®çåç§°åå»ºä¸ä¸ªå½åç©ºé´ åå»ºä¸ä¸ªç» Docker registry ä½¿ç¨ç Secret ä½¿ç¨æå®çå­å½ä»¤åå»ºä¸ä¸ª Secret åå»ºä¸ä¸ªæå®åç§°çæå¡è´¦æ· ä» kubeconfig ä¸­å é¤æå®çéç¾¤ ä» kubeconfig ä¸­å é¤æå®çä¸ä¸æ æç»ä¸ä¸ªè¯ä¹¦ç­¾åè¯·æ± æè¿°ä¸ä¸ªæå¤ä¸ªä¸ä¸æ æ¾ç¤ºå¨ kubeconfig ä¸­å®ä¹çéç¾¤ æ¾ç¤ºåå¹¶ç kubeconfig éç½®æä¸ä¸ªæå®ç kubeconfig æä»¶ æ¾ç¤ºä¸ä¸ªæå¤ä¸ªèµæº æ¸ç©ºèç¹ä»¥åå¤ç»´æ¤ ç¼è¾æå¡å¨ä¸çèµæº ç¨äº Docker éååºçé®ä»¶å°å å¨æä¸ªå®¹å¨ä¸­æ§è¡ä¸ä¸ªå½ä»¤ å°ä¸ä¸ªæå¤ä¸ªæ¬å°ç«¯å£è½¬åå°æä¸ª Pod å³äºä»»ä½å½ä»¤çå¸®å© å¦æéç©ºï¼åå°æå¡çä¼è¯äº²åæ§è®¾ç½®ä¸ºæ­¤å¼ï¼åæ³å¼ï¼'None'ã'ClientIP' å¦æéç©ºï¼ååªæå½æç»å¼æ¯å¯¹è±¡çå½åèµæºçæ¬æ¶ï¼æ³¨è§£æ´æ°æä¼æåã ä»å¨æå®åä¸ªèµæºæ¶ææã å¦æéç©ºï¼åæ ç­¾æ´æ°åªæå¨æç»å¼æ¯å¯¹è±¡çå½åèµæºçæ¬æ¶æä¼æåãä»å¨æå®åä¸ªèµæºæ¶ææã æ è®°èç¹ä¸ºå¯è°åº¦ æ è®°èç¹ä¸ºä¸å¯è°åº¦ å°ææå®çèµæºæ è®°ä¸ºå·²æå ä¿®æ¹è¯ä¹¦èµæºã ä¿®æ¹ kubeconfig æä»¶ æ­¤ä¸ºç«¯å£çåç§°æç«¯å£å·ï¼æå¡åºå°æµéå®åå°å®¹å¨ä¸çè¿ä¸ç«¯å£ãæ­¤å±æ§ä¸ºå¯éã ä»è¿åå¨æå®æ¥æ (RFC3339) ä¹åçæ¥å¿ãé»è®¤ä¸ºæææ¥å¿ãåªè½ä½¿ç¨ since-time / since ä¹ä¸ã ä¸ºæå®ç Shell(Bash æ zsh) è¾åº Shell è¡¥å¨ä»£ç ã ç¨äº Docker éååºèº«ä»½éªè¯çå¯ç  PEM ç¼ç çå¬é¥è¯ä¹¦çè·¯å¾ã ä¸ç»å®è¯ä¹¦å³èçç§é¥çè·¯å¾ã èµæºçæ¬çåææ¡ä»¶ãè¦æ±å½åèµæºçæ¬ä¸æ­¤å¼å¹éæè½è¿è¡æ©ç¼©æä½ã è¾åºå®¢æ·ç«¯åæå¡ç«¯ççæ¬ä¿¡æ¯ è¾åºææå½ä»¤çå±çº§å³ç³» æå° Pod ä¸­å®¹å¨çæ¥å¿ æ¢å¤æåçèµæº RoleBinding åºè¯¥å¼ç¨ç Role å¨éç¾¤ä¸è¿è¡ç¹å®éå è¿è¡ä¸ä¸ªæå Kubernetes API æå¡å¨çä»£ç Docker éååºçæå¡å¨ä½ç½® ä¸ºå¯¹è±¡è®¾ç½®æå®ç¹æ§ ä¸ºèµæºè®¾ç½®éæ©å¨ æ¾ç¤ºç¹å®èµæºæèµæºç»çè¯¦ç»ä¿¡æ¯ æ¾ç¤ºä¸çº¿çç¶æ --target-port çåä¹è¯ æå®å®¹å¨è¦è¿è¡çéå. å®¹å¨çéåæåç­ç¥ãå¦æçç©ºï¼è¯¥å¼å°ä¸ç±å®¢æ·ç«¯æå®ï¼ç±æå¡å¨é»è®¤è®¾ç½® æ­¤é¢ç®è¦æ±çå¯ç¨ Pod çæå°æ°éæç¾åæ¯ã æ°åå»ºçå¯¹è±¡çåç§°ã æ°åå»ºçå¯¹è±¡çåç§°ãå¦ææªæå®ï¼å°ä½¿ç¨è¾å¥èµæºçåç§°ã è¦ä½¿ç¨ç API çæå¨çåç§°ãæä¸¤ä¸ªçæå¨ã'service/v1' å 'service/v2'ãå®ä»¬ä¹é´å¯ä¸çåºå«æ¯ï¼v1 ä¸­çæå¡ç«¯å£è¢«å½åä¸º 'default'ï¼å¦æå¨ v2 ä¸­æ²¡ææå®åç§°ãé»è®¤æ¯ 'service/v2'ã è¦åå»ºçæå¡çç½ç»åè®®ãé»è®¤ä¸º âTCPâã æå¡è¦ä½¿ç¨çç«¯å£ãå¦ææ²¡ææå®ï¼åä»è¢«æ´é²çèµæºå¤å¶ è¿ä¸ªå®¹å¨çèµæºéæ±éå¶ãä¾å¦ï¼"cpu=200m,åå­=512Mi"ãè¯·æ³¨æï¼æå¡å¨ç«¯çç»ä»¶å¯è½ä¼æ ¹æ®æå¡å¨çéç½®æ¥åééå¶ï¼ä¾å¦éå¶èå´ã è¿ä¸ªå®¹å¨çèµæºéæ±è¯·æ±ãä¾å¦ï¼"cpu=200m,åå­=512Mi"ãè¯·æ³¨æï¼æå¡å¨ç«¯çç»ä»¶å¯è½ä¼æ ¹æ®æå¡å¨çéç½®æ¥åééå¶ï¼ä¾å¦éå¶èå´ã è¦åå»ºç Secret ç±»å æ¤éä¸ä¸æ¬¡çä¸çº¿ ä½¿ç¨ Pod æ¨¡æ¿æ´æ°å¯¹è±¡çèµæºè¯·æ±/éå¶ æ´æ°ä¸ä¸ªèµæºçæ³¨è§£ æ´æ°æèµæºä¸çæ ç­¾ æ´æ°ä¸ä¸ªæèå¤ä¸ªèç¹ä¸çæ±¡ç¹ ç¨äº Docker éååºèº«ä»½éªè¯çç¨æ·å æ¾ç¤ºä¸çº¿åå² å¨åªéè¾åºæä»¶ãå¦æä¸ºç©ºæ â-â åä½¿ç¨æ åè¾åºï¼å¦åå¨è¯¥ç®å½ä¸­åå»ºç®å½å±æ¬¡ç»æ åçéå¯æ å¿) kubectl æ§å¶ Kubernetes éç¾¤ç®¡çå¨ 