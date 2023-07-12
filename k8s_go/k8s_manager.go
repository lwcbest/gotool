package k8s_go

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	restclient "k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type K8SManager struct {
	namespace string
	k8sClient *kubernetes.Clientset
	k8sConfig *restclient.Config
	svc       ISvc
}

func (m *K8SManager) Init(svc ISvc) {
	m.namespace = "default"
	m.initLocalConfig()
	m.initK8SClient()
	m.svc = svc
}

func (m *K8SManager) StopSvc() {
	m.svc.InitMetaData()
	fmt.Println("***| start get resource |***")
	services := m.getCurrentSvc()
	dms := m.getCurrentDeployment()
	configs := m.getCurrentConfigMap()

	fmt.Println("***| start del resource |***")
	m.delSvc(m.svc.GetName(), services)
	m.delDeploy(m.svc.GetName(), dms)
	m.delConfig(m.svc.GetName()+"-config", configs)
}

func (m *K8SManager) StartSvc() {
	m.StopSvc()
	fmt.Println("***| start create resource |***")
	m.startConfig(m.svc.GetConfigMapMeta())
	m.startSvc(m.svc.GetServiceMeta())
	m.startDeploy(m.svc.GetDeploymentMeta())
}

func (m *K8SManager) initLocalConfig() {
	// 解析命令行参数获取kubeconfig文件路径
	kubeConfig := flag.String("kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	flag.Parse()
	// 使用kubeconfig文件创建一个Kubernetes客户端
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConfig)
	if err != nil {
		panic(err.Error())
	}
	m.k8sConfig = config
}

func (m *K8SManager) initK8SClient() {
	clientset, err := kubernetes.NewForConfig(m.k8sConfig)
	if err != nil {
		panic(err.Error())
	}

	m.k8sClient = clientset
}

func (m *K8SManager) getCurrentSvc() *corev1.ServiceList {
	services, err := m.k8sClient.CoreV1().Services(m.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, service := range services.Items {
		fmt.Printf("Service Name: %s\n", service.Name)
		for _, port := range service.Spec.Ports {
			fmt.Printf("\tPort Name: %s, Port Number: %d\n", port.Name, port.Port)
		}
	}
	return services
}

func (m *K8SManager) getCurrentDeployment() *appsv1.DeploymentList {
	dms, err := m.k8sClient.AppsV1().Deployments(m.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, dm := range dms.Items {
		fmt.Printf("DeployMent Name: %s\n", dm.Name)
		for _, container := range dm.Spec.Template.Spec.Containers {
			fmt.Printf("\tPort Name: %s, Port Numbers: %v\n", container.Name, container.Ports)
		}
	}
	return dms
}

func (m *K8SManager) getCurrentConfigMap() *corev1.ConfigMapList {
	configMaps, err := m.k8sClient.CoreV1().ConfigMaps(m.namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, config := range configMaps.Items {
		fmt.Printf("config Name: %s\n", config.Name)
		for key, value := range config.Data {
			fmt.Printf("\tKey Name: %s, Value: %v\n", key, value)
		}
	}
	return configMaps
}

func (m *K8SManager) delSvc(name string, services *corev1.ServiceList) {
	for _, service := range services.Items {
		if name == service.Name {
			err := m.k8sClient.CoreV1().Services(m.namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Printf("\tDel Svc Fail....  %s \n", err)
			}
			fmt.Printf("\tDel Svc Name: %s \n", name)
		}
	}
}

func (m *K8SManager) delDeploy(name string, dms *appsv1.DeploymentList) {
	for _, deploy := range dms.Items {
		if name == deploy.Name {
			err := m.k8sClient.AppsV1().Deployments(m.namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Printf("\tDel deploy Fail....  %s \n", err)
			}
			fmt.Printf("\tDel deploy Name: %s \n", name)
		}
	}
}

func (m *K8SManager) delConfig(name string, configs *corev1.ConfigMapList) {
	for _, config := range configs.Items {
		if name == config.Name {
			err := m.k8sClient.CoreV1().ConfigMaps(m.namespace).Delete(context.TODO(), name, metav1.DeleteOptions{})
			if err != nil {
				fmt.Printf("\tDel config Fail....  %s \n", err)
			}
			fmt.Printf("\tDel config Name: %s \n", name)
		}
	}
}

func (m *K8SManager) startSvc(svc *corev1.Service) {
	// 创建Redis服务
	service, err := m.k8sClient.CoreV1().Services(m.namespace).Create(context.TODO(), svc, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\tCreated service %s\n", service.GetName())
}

func (m *K8SManager) startDeploy(dm *appsv1.Deployment) {
	// 创建Redis服务
	dm, err := m.k8sClient.AppsV1().Deployments(m.namespace).Create(context.TODO(), dm, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\tCreated deploy %s\n", dm.GetName())
}

func (m *K8SManager) startConfig(config *corev1.ConfigMap) {
	// 创建Redis服务
	config, err := m.k8sClient.CoreV1().ConfigMaps(m.namespace).Create(context.TODO(), config, metav1.CreateOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("\tCreated config %s\n", config.GetName())
}
