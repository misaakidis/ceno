package plugins.CENO.Client;

import plugins.CENO.CENOErrCode;
import plugins.CENO.CENOException;
import freenet.pluginmanager.PluginHTTPException;
import freenet.support.api.HTTPRequest;

public class ClientHandler extends AbstractCENOClientHandler {

	private static final String pluginPath = "/plugins/" + CENOClient.class.getName();
	private static final LookupHandler lookupHandler = new LookupHandler();

	public String handleHTTPGet(HTTPRequest request) throws PluginHTTPException {
		String path = request.getPath().replaceFirst(pluginPath, "");
		if (path.isEmpty() || path.equals("/") || path.equals("/index.html")) {
			return printStaticHTML("Resources/index.html");
		} else if (path.startsWith("/lookup")) {
			return lookupHandler.handleHTTPGet(request);
		}
		if (isClientHtml(request)) {
			return "404: Requested path is invalid.";
		} else {
			return returnErrorJSON(new CENOException(CENOErrCode.LCS_HANDLER_URL_INVALID));
		}
	}

	public String handleHTTPPost(HTTPRequest request) throws PluginHTTPException {
		String path = request.getPath().replaceFirst(pluginPath, "");
		if (!path.isEmpty() && path.startsWith("/fetch")) {
			RequestSender.requestFromBridge(request.getParam("url", ""));
			return "Sent passive request";
		}
		return "404: Requested path is invalid.";
	}

}
