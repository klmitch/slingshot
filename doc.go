// Package slingshot contains a plugin manager built on top of the
// standard plugin package.  Plugins are loaded using the Load
// function, which takes the filesystem path to the compiled plugin
// and a map of parameters.  The Load function uses the plugin package
// to load the plugin and invokes its SlingshotInit function, which
// must be declared to take as arguments a Slingshot interface object
// and the parameters (declared as a map from string to generic
// interface).  The SlingshotInit function is then expected to use the
// Slingshot object's Register method, passing it a namespace, a key,
// the plugin object itself, and zero or more plugin options.
//
// The slingshot package divides plugins up into namespaces.  The
// namespaces should be unique for the application, e.g.,
// "github.com/klmitch/slingshot".  If the application requires
// multiple namespaces for different types of plugins, those
// namespaces should have a common prefix, e.g.,
// "github.com/klmitch/slingshot/hooks" and
// "github.com/klmitch/slingshot/drivers".
//
// Plugins are further divided up by a "key"; this is again a simple
// string, and its use will depend on how the application uses the
// plugin.  For instance, in driver-style plugins, the key identifies
// the driver, while in hook-style plugins, the key names the hook to
// be triggered.
//
// Plugins come in many varieties.  The simplest to understand is
// perhaps the driver pattern plugin, where a driver is specified by a
// name.  The desired driver may then be retrieved by simply calling
// GetPlugin with the appropriate namespace and driver name; this
// retrieves only the first-registered plugin with the given name.
// Hook-pattern plugins, on the other hand, expect to have all plugins
// with the same key invoked whenever the hook is triggered; to
// implement this pattern, call GetAllPlugins with the appropriate
// namespace and hook name, then iterate over the returned list.
//
// A more complex pattern is the extension pattern.  In this pattern,
// again, all the plugins are invoked in order, but each plugin calls
// the next plugin, and has the opportunity to process its return
// value.  To make this easier to accommodate, the slingshot package
// provides the IterPlugins utility function, which encapsulates the
// list of plugins in an object with a Next method returning the next
// plugin.  With this, an extension function could be written like so:
//
//     func Extension(nextPlug *PluginIter, finalFunc func() error) {
//         next := nextPlug.Next()
//         if next == nil {
//             return finalFunc()
//         }
//
//         return callExtension(next, nextPlug, finalFunc)
//     }
//
// In this example, callExtension is assumed to be a function that
// calls the plugin function; more on that below.
//
// Some applications may have built-in plugins.  For instance, a
// driver-style plugin may provide a mock driver.  Such plugins can be
// registered directly using the Register function, which has the same
// calling convention as the Register method of the Slingshot object
// passed to the plugin initialization function.
//
// What is a plugin?  The GetPlugin and GetAllPlugins functions
// ultimately return a PluginMeta object.  This object contains more
// than just the plugin itself; it contains a large amount of
// metadata, such as the filesystem path to the plugin, the plugin's
// base filename, and the namespace and key of the plugin.  Additional
// options can be passed to the Register function to set a name, a
// plugin version, a license description, or documentation.  An
// integer APIVersion can also be specified, which could be used to
// allow an application to still function with plugins implementing an
// older version of the plugin's interface.  Finally, any arbitrary
// metadata can be provided, which could be used by the application
// for additional complex filtering when iterating over the list of
// plugins.
//
// The Plugin element of the PluginMeta object is the actual plugin
// passed to the Register function.  This is declared as a generic
// interface, so this may be anything, from a simple string to a
// function to an object implementing an interface.  The callExtension
// function described above could thus be something of the form:
//
//     func callExtension(plug *PluginMeta, nextPlug *PluginIter, finalFunc func() error) {
//         plug.Plugin.(func(*PluginIter, func() error) error)(nextPlug, finalFunc)
//     }
//
// Finally, the slingshot package contains full-featured mocks, built
// on the "github.com/stretchr/testify/mock" mocking package.  The
// MockRegistry type allows mocking the slingshot plugin registry.
// This mock can be installed for the use of functions such as
// GetPlugin by using SetRegistry from a test function like so:
//
//     defer SetRegistry(SetRegistry(mockReg))
//
// The MockSlingshot type can be passed to a plugin initialization
// function.  Finally, there is a MockNamespace type, which has a Name
// element to set the return result of the Namespace method call.
//
// For more information on using the mocks, check out the
// documentation for "github.com/stretchr/testify/mock".
package slingshot
